// Package requests implements methods for sending HTTP requests to Gladius server.
package requests

import (
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
