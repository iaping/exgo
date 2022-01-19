package okex

import (
	"github.com/iaping/exgo/okex/api/asset"
)

// https://www.ouyi.fit/docs-v5/zh/#rest-api-funding-get-currencies
func (c *Client) AssetCurrencies() ([]*asset.Currencies, error) {
	req := asset.NewCurrenciesRequest()
	resp := asset.NewCurrenciesResponse()
	if err := c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// no test
// https://www.ouyi.fit/docs-v5/zh/#rest-api-funding-get-balance
func (c *Client) AssetBalances(ccy string) ([]*asset.Balances, error) {
	req := asset.NewBalancesRequest(ccy)
	resp := asset.NewBalancesResponse()
	if err := c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// https://www.ouyi.fit/docs-v5/zh/#rest-api-funding-asset-bills-details
func (c *Client) AssetBills(ccy string, typed, limit int, after, before int64) ([]*asset.Bills, error) {
	req := asset.NewBillsRequest(ccy, typed, limit, after, before)
	resp := asset.NewBillsResponse()
	if err := c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
