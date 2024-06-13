package requests

type RawOrders []byte
type OrderStatus string

const (
	OPEN    OrderStatus = "open"
	EXPIRED OrderStatus = "expired"
	FILLED  OrderStatus = "filled"
)

// A representation of a query-string for a
// GET request to the '/orders' endpoint.
type OrdersGET struct {
	// If empty URL from constants will be used.
	GladiusUrl string
	Status     OrderStatus
	Options    *OrdersOpts
}

// Optional query params that can be used to specify
// criterias of order's search more accurately.
// There are server rules that don't allow certain fields
// to be presented with other specific fields or without other fields:
// TODO: implement validation of these rules (?)
type OrdersOpts struct {
	ChainID int

	SellToken string
	BuyToken  string

	// Querying with both 'swapper' and 'chainId'
	// is not currently supported.
	// Therefore, if 'swapper' is set 'chainID' will be omitted.
	// TODO: throw error in such case.
	Swapper string

	// A maximum number of orders to return.
	Limit int

	// Return orders with sepcific hashes.
	// If one of the hashes in a slice is invalid,
	// then only orders woth valid hashes will be returned.
	OrderHashes []string

	SortKeys *SortKeys

	// *Optional* API key
	GladiusAuthKey string
}

// Query-params for sorting of orders.
type SortKeys struct {
	// Sort query used with 'createdAt' sort key.
	// *Shouldn't* include 'createdAt' part!
	// https://api.uniswap.org/v2/uniswapx/docs ('/orders&sort')
	Sort string
	// Boolean to sort query results by descending sort key.
	Desc bool
}
