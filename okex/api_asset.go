package okex

import (
	"github.com/iaping/exgo/okex/api/asset"
)

// https://www.ouyi.fit/docs-v5/zh/#rest-api-funding-get-currencies
func (c *Client) AssetCurrencies() ([]*asset.Currencies, error) {
	req := asset.NewCurrenciesRequest()
	var resp asset.CurrenciesResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// no test
// https://www.ouyi.fit/docs-v5/zh/#rest-api-funding-get-balance
func (c *Client) AssetBalances(ccy string) ([]*asset.Balances, error) {
	req := asset.NewBalancesRequest(ccy)
	var resp asset.BalancesResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// https://www.ouyi.fit/docs-v5/zh/#rest-api-funding-asset-bills-details
func (c *Client) AssetBills(ccy string, typed, limit int, after, before int64) ([]*asset.Bills, error) {
	req := asset.NewBillsRequest(ccy, typed, limit, after, before)
	var resp asset.BillsResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
