// This file contains functions for converting 'GladiusOrder' and its JSON
// analogues into ResolvedOrders, with real-time
package gladiusOrder

import (
	"github.com/RubiconDeFi/gladius-go-sdk/binds"
	"github.com/RubiconDeFi/gladius-go-sdk/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"sync"
)

// Resolves a slice of gladius orders, received from server
// into a slice of 'ResolvedOrder', that contains easier-to-parse-through
// values of underlying JSON orders, and Point-In-Time values for
// decaying input and outputs, resolved based on the value
// of 'time.Now().Unix()' on the function call.
//
// 'bopt' argument allows to add fee structs into the 'ResolvedOrder.Outputs',
// by adding 'EthClient' and either specifying an 'Address' of the target
// fee controller or by specifying 'ChainID' to use constant value.
// Note, if specified it will execute RPC call to get fees.
// Set it to 'nil' if fee inclusion isn't needed.
func (g *GladiusOrdersJSON) Resolve(bopt *binds.BindOpts) ([]*contracts.ResolvedOrder, error) {
	orders := g.Orders
	r := make([]*contracts.ResolvedOrder, len(orders))
	errch := make(chan error, len(orders))

	var wg sync.WaitGroup
	wg.Add(len(orders))

	for i := 0; i < len(orders); i++ {
		go func(i int) {
			ri, err := orders[i].Resolve(bopt)
			if err != nil {
				errch <- err
				return
			}
			r[i] = ri

			wg.Done()
		}(i)
	}
	wg.Wait()
	close(errch)

	for err := range errch {
		if err != nil {
			return nil, err
		}
	}

	return r, nil
}

// Resolves individual gladius order into individual 'ResolvedOrder'.
// The same rules apply to this function as for 'Resolve' for a slice of orders.
func (g *GladiusOrderJSON) Resolve(bopt *binds.BindOpts) (*contracts.ResolvedOrder, error) {
	rout := make([]contracts.OutputToken, len(g.Outputs))

	dStart := ParseDecayStartTime(g.EncodedOrder)
	dEnd := ParseDecayEndTime(g.EncodedOrder)

	var wg sync.WaitGroup
	wg.Add(len(rout))

	for i := 0; i < len(rout); i++ {
		go func(i int) {
			rout[i] = g.Outputs[i].Resolve(dStart, dEnd)
			wg.Done()
		}(i)
	}
	wg.Wait()

	r := &contracts.ResolvedOrder{
		Info:    g.OrderInfo(),
		Input:   g.Input.Resolve(dStart, dEnd),
		Outputs: rout,
		Sig:     hexutil.MustDecode(g.Signature),
		Hash:    [32]byte(hexutil.MustDecode(g.OrderHash)),
	}

	if bopt != nil {
		out, err := bopt.GetFee(r)

		if err != nil {
			return r, err
		}
		r.Outputs = append(r.Outputs, out...)
	}

	return r, nil
}

// Extracts 'OrderInfo' from
func (g *GladiusOrderJSON) OrderInfo() contracts.OrderInfo {
	eo := g.EncodedOrder

	return contracts.OrderInfo{
		Reactor:                      ParseReactor(eo),
		Swapper:                      ParseSwapper(eo),
		Nonce:                        ParseNonce(eo),
		Deadline:                     ParseDeadline(eo),
		AdditionalValidationContract: ParseAdditionalValidationContract(eo),
		// TODO: returns []byte{} for now.
		AdditionalValidationData: ParseAdditionalValidationData(eo),
	}
}

func (o DutchOutput) Resolve(dStart, dEnd *big.Int) contracts.OutputToken {
	return contracts.OutputToken{
		Token:     common.HexToAddress(o.Token),
		Amount:    o.Decay(dStart, dEnd),
		Recipient: common.HexToAddress(o.Recipient),
	}
}

func (i DutchInput) Resolve(dStart, dEnd *big.Int) contracts.InputToken {
	endamt := new(big.Int)
	endamt.SetString(i.EndAmount, 10)

	return contracts.InputToken{
		Token:     common.HexToAddress(i.Token),
		Amount:    i.Decay(dStart, dEnd),
		MaxAmount: endamt,
	}
}
