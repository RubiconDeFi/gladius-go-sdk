package requests

import (
	"errors"
	"github.com/RubiconDeFi/gladius-go-sdk/constants"
	"github.com/valyala/fasthttp"
	"strconv"
)

func sendGET(url string, authKey string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod("GET")
	req.SetRequestURI(url)

	if authKey != "" {
		req.Header.Set("x-api-key", authKey)
	}

	resp := fasthttp.AcquireResponse()

	err := fasthttp.Do(req, resp)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		code := strconv.Itoa(resp.StatusCode())

		return nil, errors.New(
			"Request failed with status code:" + code)
	}

	body := make([]byte, len(resp.Body()))
	copy(body, resp.Body())
	fasthttp.ReleaseResponse(resp)

	return body, nil
}

// Builds a query-string for a GET '/orders' request.
func (req *OrdersGET) ToUrl() (url string) {
	if req.GladiusUrl != "" {
		url = req.GladiusUrl + "orders?"
	} else {
		url = constants.GLADIUS_URL + "orders?"
	}

	if req.Status != "" {
		url += "&orderStatus=" + string(req.Status)
	}

	opts := req.Options

	if opts != nil {
		if opts.ChainID != 0 {
			url += "&chainId=" + strconv.Itoa(opts.ChainID)
		}

		// TODO: handle incorrect params.
		if opts.SellToken != "" {
			url += "&sellToken=" + opts.SellToken
		}

		if opts.BuyToken != "" {
			url += "&buyToken=" + opts.BuyToken
		}

		if opts.Swapper != "" {
			url += "&swapper=" + opts.Swapper
		}

		if opts.Limit != 0 {
			url += "&limit=" + strconv.Itoa(opts.Limit)
		}

		if opts.OrderHashes != nil {
			h := opts.OrderHashes
			url += "&orderHashes="

			for i := 0; i < len(h); i++ {
				url += h[i]

				if i+1 != len(h) {
					url += ","
				}
			}
		}

		sk := opts.SortKeys
		if sk != nil {
			if sk.Sort != "" {
				url += "&sortKey=createdAt&sort=" + sk.Sort
			}

			if sk.Desc {
				url += "&desc=true"
			}
		}
	}

	return
}
