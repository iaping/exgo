package okex

import (
	"github.com/iaping/exgo/okex/api/market"
)

// https://www.okx.com/docs-v5/zh/#rest-api-market-data-get-candlesticks
func (c *Client) MarketCandles(param *market.CandlesParam) ([][]string, error) {
	req := market.NewCandlesRequest(param)
	resp := market.NewCandlesResponse()

	if err := c.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
