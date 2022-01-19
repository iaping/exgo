package okex

import (
	"github.com/iaping/exgo/okex/api/account"
	"github.com/iaping/exgo/okex/api/subaccount"
)

// https://www.ouyi.fit/docs-v5/zh/#rest-api-subaccount-get-sub-account-balance
func (c *Client) SubaccountBalances(subAcct string) ([]*account.Balance, error) {
	req := subaccount.NewSubaccountBalancesRequest(subAcct)
	resp := account.NewBalanceResponse()
	if err := c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
