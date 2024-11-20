package requests

import (
	"github.com/RubiconDeFi/gladius-go-sdk/binds"
	"github.com/RubiconDeFi/gladius-go-sdk/constants"
	glo "github.com/RubiconDeFi/gladius-go-sdk/order/gladiusOrder"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

const (
	SELL = "0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1"
	BUY  = "0x7F5c764cBc14f9669B88837ca1490cCa17c31607"
)

var (
	req = &OrdersGET{
		// GladiusUrl: default
		Status: EXPIRED,
		Options: &OrdersOpts{
			ChainID:   constants.OPTIMISM_MAINNET,
			SellToken: SELL,
			BuyToken:  BUY,
			Limit:     228,
		},
	}
	statuses [3]OrderStatus = [3]OrderStatus{OPEN, EXPIRED, FILLED}
	rpc                     = "https://mainnet.optimism.io/"
)

func TestGetRawOrders(t *testing.T) {
	raw, err := GetRawOrders(req)

	if err != nil {
		t.Error(err)
	}

	if raw == nil {
		t.Error("raw is nil")
	}

	if len(raw) == 0 {
		t.Error("got no raw orders")
	}
}

func TestGetGladiusOrders(t *testing.T) {
	o, err := GetGladiusOrders(req)

	if err != nil {
		t.Error(err)
	}

	if len(o.Orders) > 0 {
		validateStatus(req.Status, o, t)
	}
}

func TestGetResolvedGldOrders(t *testing.T) {
	ethc, err := ethclient.Dial(rpc)
	if err != nil {
		t.Error(err)
	}
	bopt := &binds.BindOpts{
		ChainID:   10,
		EthClient: ethc,
	}
	o, err := GetResolvedGladiusOrders(req, bopt)

	if err != nil {
		t.Error(err)
	}

	// Check that all the order hashes are unique
	var usedHashes map[[32]byte]bool = make(map[[32]byte]bool)
	for i := 0; i < len(o); i++ {
		if usedHashes[o[i].Hash] {
			t.Error("hash was already used!")
		}
		usedHashes[o[i].Hash] = true
	}

	if len(o) == 0 {
		t.Error("no orders")
	}
}

func TestGetGladiusOrdersWithStatus(t *testing.T) {
	for i := 0; i < len(statuses); i++ {
		s := statuses[i]
		req.Status = s

		o, err := GetGladiusOrders(req)

		if err != nil {
			t.Error(err)
		}

		if len(o.Orders) > 0 {
			validateStatus(req.Status, o, t)
		}

	}
}

func validateStatus(s OrderStatus, o *glo.GladiusOrdersJSON, t *testing.T) {
	orders := o.Orders

	for i := 0; i < len(orders); i++ {
		ss := orders[i].OrderStatus

		if ss != string(s) {
			t.Errorf("Invalid status\nhave: [%v]\nwant: [%v]\n", ss, s)
		}
	}
}
