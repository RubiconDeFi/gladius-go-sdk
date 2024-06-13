package requests

import (
	"testing"
)

var ezr = &OrdersGET{
	Status: OPEN,
}

var fullr = &OrdersGET{
	GladiusUrl: "http://yo.bro/swedish-auction/",
	Status:     OPEN,
	Options: &OrdersOpts{
		ChainID:     322,
		SellToken:   SELL,
		BuyToken:    BUY,
		Swapper:     "0x1337",
		Limit:       228,
		OrderHashes: []string{"a", "z", "g"},
		SortKeys:    &SortKeys{"gt(1)", true},
	},
}

// this url won't work server-side tho.
const (
	F_URL = "http://yo.bro/swedish-auction/orders?&orderStatus=open&chainId=322&sellToken=0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1&buyToken=0x7F5c764cBc14f9669B88837ca1490cCa17c31607&swapper=0x1337&limit=228&orderHashes=a,z,g&sortKey=createdAt&sort=gt(1)&desc=true"
	E_URL = "https://gladius.rubicon.finance/dutch-auction/orders?&orderStatus=open"
)

func TestFullUrl(t *testing.T) {
	url := fullr.ToUrl()

	if url != F_URL {
		t.Errorf("\nInvalid URL\nhave: [%v]\nwant: [%v]\n", url, F_URL)
	}
}

func TestEzUrl(t *testing.T) {
	url := ezr.ToUrl()

	if url != E_URL {
		t.Errorf("\nInvalid URL\nhave: [%v]\nwant: [%v]\n", url, E_URL)
	}
}
