// This file contains target struct for a JSON response
// from a gladius server, that should be used
// for JSON unmarshalling.
// And also reusable 'Dutch[I/O]'.
package gladiusOrder

// Target struct for a server's response.
type GladiusOrdersJSON struct {
	Orders []GladiusOrderJSON `json:"orders"`
	Cursor string             `json:"cursor"`
}

// Struct, that represent individual order.
type GladiusOrderJSON struct {
	Outputs      []DutchOutput `json:"outputs"`
	OrderStatus  string        `json:"orderStatus"`
	CreatedAt    int           `json:"createdAt"`
	EncodedOrder string        `json:"encodedOrder"`
	Signature    string        `json:"signature"`
	ChainID      int           `json:"chainId"`
	OrderHash    string        `json:"orderHash"`
	Input        DutchInput    `json:"input"`
	Price        float64       `json:"price"`
	Type         string        `json:"type"`
}

// Decaying input.
type DutchInput struct {
	EndAmount   string `json:"endAmount"`
	Token       string `json:"token"`
	StartAmount string `json:"startAmount"`
}

// Decaying output.
type DutchOutput struct {
	Recipient   string `json:"recipient"`
	StartAmount string `json:"startAmount"`
	EndAmount   string `json:"endAmount"`
	Token       string `json:"token"`
}
