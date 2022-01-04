package okex

import "github.com/iaping/exgo/okex/api/account"

// https://www.ouyi.fit/docs-v5/zh/#rest-api-account-get-balance
func (c *Client) AccountBalance(ccy string) ([]*account.Balance, error) {
	req := account.NewBalanceRequest(ccy)
	var resp account.BalanceResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
