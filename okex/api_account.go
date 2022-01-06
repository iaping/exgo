package okex

import (
	"github.com/iaping/exgo/okex/api/account"
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

// https://www.ouyi.fit/docs-v5/zh/#rest-api-account-get-bills-details-last-7-days
func (c *Client) AccountBills(instType, ccy, mgnMode, after, before string, ctType, tp, subType, limit int) ([]*account.Bills, error) {
	req := account.NewBillsRequest(instType, ccy, mgnMode, after, before, ctType, tp, subType, limit)
	var resp account.BillsResponse
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
