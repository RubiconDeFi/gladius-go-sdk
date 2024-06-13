package requests

import (
	"bytes"
	"encoding/json"
	glo "github.com/RubiconDeFi/gladius-go-sdk/order/gladiusOrder"
)

func (r RawOrders) ToGladiusOrders() (*glo.GladiusOrdersJSON, error) {
	o := &glo.GladiusOrdersJSON{}
	err := json.NewDecoder(bytes.NewReader(r)).Decode(o)

	return o, err
}
