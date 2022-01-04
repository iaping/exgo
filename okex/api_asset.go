package okex

import "github.com/iaping/exgo/okex/api/asset"

// https://www.ouyi.fit/docs-v5/zh/#rest-api-funding-get-currencies
func (c *Client) AssetCurrencies() ([]*asset.Currencies, error) {
	req := asset.NewCurrenciesRequest()
	var resp asset.CurrenciesResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
