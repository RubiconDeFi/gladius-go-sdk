// Package requests implements methods for sending HTTP requests to Gladius server.
package requests

import (
	"github.com/RubiconDeFi/gladius-go-sdk/binds"
	"github.com/RubiconDeFi/gladius-go-sdk/contracts"
	glo "github.com/RubiconDeFi/gladius-go-sdk/order/gladiusOrder"
)

// Returns raw byte response from the server.
func GetRawOrders(req *OrdersGET) (raw RawOrders, err error) {
	raw, err = sendGET(req.ToUrl(), req.Options.GladiusAuthKey)
	return
}

// Returns decoded response from the server.
func GetGladiusOrders(req *OrdersGET) (*glo.GladiusOrdersJSON, error) {
	raw, err := GetRawOrders(req)

	if err != nil {
		return nil, err
	}

	o, err := raw.ToGladiusOrders()

	return o, err
}

// Gets gladius orders from the server and resolves them.
func GetResolvedGladiusOrders(req *OrdersGET, feeOpts *binds.BindOpts) ([]*contracts.ResolvedOrder, error) {
	raw, err := GetRawOrders(req)
	if err != nil {
		return nil, err
	}

	o, err := raw.ToGladiusOrders()
	if err != nil {
		return nil, err
	}

	return o.Resolve(feeOpts)
}
