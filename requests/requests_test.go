package requests

import (
	"github.com/RubiconDeFi/gladius-go-sdk/constants"
	glo "github.com/RubiconDeFi/gladius-go-sdk/order/gladiusOrder"
	"testing"
)

const (
	SELL = "0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1"
	BUY  = "0x7F5c764cBc14f9669B88837ca1490cCa17c31607"
)

var req = &OrdersGET{
	// GladiusUrl: default
	Status: EXPIRED,
	Options: &OrdersOpts{
		ChainID:   constants.OPTIMISM_MAINNET,
		SellToken: SELL,
		BuyToken:  BUY,
		Limit:     228,
	},
}
var statuses [3]OrderStatus = [3]OrderStatus{OPEN, EXPIRED, FILLED}

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
