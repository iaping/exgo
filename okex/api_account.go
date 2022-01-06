package okex

import (
	"github.com/iaping/exgo/okex/api/account"
	"github.com/iaping/exgo/okex/api/account/subaccount"
)

// https://www.ouyi.fit/docs-v5/zh/#rest-api-account-get-balance
func (c *Client) AccountBalance(ccy string) ([]*account.Balance, error) {
	req := account.NewBalanceRequest(ccy)
	var resp account.BalanceResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// https://www.ouyi.fit/docs-v5/zh/#rest-api-account-get-account-configuration
func (c *Client) AccountConfig() ([]*account.Config, error) {
	req := account.NewConfigRequest()
	var resp account.ConfigResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// https://www.ouyi.fit/docs-v5/zh/#rest-api-subaccount-get-sub-account-balance
func (c *Client) AccountSubaccountBalances(subAcct string) ([]*account.Balance, error) {
	req := subaccount.NewSubaccountBalancesRequest(subAcct)
	var resp account.BalanceResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
