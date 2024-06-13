package gladiusOrder

import (
	"github.com/RubiconDeFi/gladius-go-sdk/binds"
	"github.com/RubiconDeFi/gladius-go-sdk/constants"
	"github.com/RubiconDeFi/gladius-go-sdk/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"testing"
	"time"
)

const (
	START_AMT  = "100000000000000000000"
	END_AMT    = "200000000000000000000"
	HALF_DECAY = "150000000000000000000"
	SELL       = "0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1"
	BUY        = "0x7F5c764cBc14f9669B88837ca1490cCa17c31607"
	RUBI_RECIP = "0x752748DeaF25cf58b60d4C4209d7F200AeE4Ef14"
)

var (
	DutchIn = DutchInput{
		Token: SELL, StartAmount: START_AMT,
		EndAmount: END_AMT}
	DutchOut = DutchOutput{
		Token: BUY, StartAmount: END_AMT,
		EndAmount: START_AMT, Recipient: BUY}
	Gorder = GladiusOrderJSON{
		Outputs: []DutchOutput{{
			Recipient:   "0x4476a2c123c408ce0caf1754969e7ba78e7966d8",
			StartAmount: "399980000000000",
			EndAmount:   "399980000000000",
			Token:       "0x7F5c764cBc14f9669B88837ca1490cCa17c31607"}},
		OrderStatus:  "expired",
		CreatedAt:    1716672996,
		EncodedOrder: "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000665259d80000000000000000000000000000000000000000000000000000000066525d5b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000da10009cbd5d07dd0cecc66161fc93d7c9000da10000000000000000000000000000000000000000000000d8d726b7177a8000000000000000000000000000000000000000000000000000d8d726b7177a80000000000000000000000000000000000000000000000000000000000000000002200000000000000000000000000000000000000000000000d8d726b7177a80000000000000000000000000000098169248bdf25e0e297ea478ab46ac24058fac780000000000000000000000004476a2c123c408ce0caf1754969e7ba78e7966d800000000000000000000000000000000000000000000000000000000000001b50000000000000000000000000000000000000000000000000000000066525d5c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000007f5c764cbc14f9669b88837ca1490cca17c3160700000000000000000000000000000000000000000000000000016bc799d1380000000000000000000000000000000000000000000000000000016bc799d138000000000000000000000000004476a2c123c408ce0caf1754969e7ba78e7966d8",
		Signature:    "0x67d1c7d7ba5375532c4b5ad610745471d9a4816ea6859a08f0adc3164b554ca672fc037058c4d3f98f6ca76da9cf3220da061ee8a6fa3acfb91a71f85e66dfde1b",
		ChainID:      10,
		OrderHash:    "0x87f1a1fb51b3c609d40abcbd8f52700b7a0a5e0ad3e44b723f333e8cf50182c6",
		Input: DutchInput{
			EndAmount:   "4000000000000000000000",
			Token:       "0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1",
			StartAmount: "4000000000000000000000",
		},
		Price: 1.1, // Doesn't really matter here
		Type:  "Dutch"}

	Rpc = "https://rpc.ankr.com/optimism"
)

func TestDecay(t *testing.T) {
	now := time.Now().Unix()

	// Decay already started
	decayStart := big.NewInt(now - 10)
	decayEnd := big.NewInt(now + 10)

	halfDecay := new(big.Int)
	halfDecay.SetString(HALF_DECAY, 10)

	in := DutchIn.Decay(decayStart, decayEnd)
	if in.Cmp(halfDecay) != 0 {
		t.Errorf("Invalid decay\nhave: [%v] \nwant: [%v]", in, halfDecay)
	}

	out := DutchOut.Decay(decayStart, decayEnd)
	if out.Cmp(halfDecay) != 0 {
		t.Errorf("Invalid decay\nhave: [%v] \nwant: [%v]", out, halfDecay)
	}
}

func TestResolveGorder(t *testing.T) {
	r, err := Gorder.Resolve(nil)

	if err != nil {
		t.Error(err)
	}

	//~~~~~~ Validate 'OrderInfo'
	if r.Info.Reactor.Hex() != constants.REACTOR_ADDRESS_MAPPING[10]["Dutch"] {
		t.Errorf("Invalid reactor\nhave: [%v] \nwant: [%v]",
			r.Info.Reactor,
			constants.REACTOR_ADDRESS_MAPPING[10]["Dutch"])
	}
	if strings.ToLower(r.Info.Swapper.Hex()) != strings.ToLower(Gorder.Outputs[0].Recipient) {
		t.Errorf("Invalid swapper\nhave: [%v] \nwant: [%v]",
			r.Info.Swapper.Hex(),
			Gorder.Outputs[0].Recipient)
	}
	if r.Info.Nonce.Cmp(big.NewInt(437)) != 0 {
		t.Errorf("Invalid nonce\nhave: [%v] \nwant: [%v]", r.Info.Nonce, 437)
	}
	if r.Info.Deadline.Cmp(big.NewInt(1716673884)) != 0 {
		t.Errorf("Invalid deadline\nhave: [%v] \nwant: [%v]", r.Info.Deadline, 1716673884)
	}
	if r.Info.AdditionalValidationContract != constants.ZERO_ADDRESS {
		t.Errorf("Invalid deadline\nhave: [%v] \nwant: [%v]",
			r.Info.AdditionalValidationContract.Hex(),
			constants.ZERO_ADDRESS)
	}
	if len(r.Info.AdditionalValidationData) != 0 {
		t.Errorf("Invalid deadline\nhave: [%v] \nwant: [%v]",
			r.Info.AdditionalValidationData,
			[]byte{})
	}

	inendamt := new(big.Int)
	inendamt.SetString(Gorder.Input.EndAmount, 10)

	//~~~~~~ Validate 'Input'
	if r.Input.Token.Hex() != Gorder.Input.Token {
		t.Errorf("Invalid input.token\nhave: [%v] \nwant: [%v]",
			r.Input.Token.Hex(),
			Gorder.Input.Token)
	}
	// Will be equal to 'endAmount' because order has already expired.
	if r.Input.Amount.Cmp(inendamt) != 0 && r.Input.MaxAmount.Cmp(inendamt) != 0 {
		t.Errorf("Invalid input.amount\nhave: [%v] \nwant: [%v]",
			r.Input.Token.Hex(),
			Gorder.Input.Token)
	}

	validateOutputs(r.Outputs, t)
}

func TestResolveWithFee(t *testing.T) {
	ethc, err := ethclient.Dial(Rpc)
	if err != nil {
		t.Error(err)
	}

	bindOpts := &binds.BindOpts{
		EthClient: ethc,
		ChainID:   Gorder.ChainID,
	}
	r, err := Gorder.Resolve(bindOpts)
	if err != nil {
		t.Error(err)
	}

	if len(r.Outputs) != 2 {
		t.Error("No fee")
	}

	validateOutputs(r.Outputs, t)
}

func validateOutputs(o []contracts.OutputToken, t *testing.T) {
	if len(o) == 0 {
		t.Error("no 'Outputs' in 'ResolvedOrder'")
	}

	for i := 0; i < len(o); i++ {
		if o[i].Token.Hex() != Gorder.Outputs[0].Token {
			t.Errorf("Invalid output.token\nhave: [%v] \nwant: [%v]",
				o[i].Token.Hex(),
				Gorder.Outputs[0].Token)
		}

		if strings.ToLower(o[i].Recipient.Hex()) != Gorder.Outputs[0].Recipient {
			// Last iteration
			if i+1 == len(o) {
				if o[i].Recipient.Hex() != RUBI_RECIP {
					t.Error("invalid recip in fee")
				}
			} else {
				t.Error("invalid recip")
			}
		}
	}
}
