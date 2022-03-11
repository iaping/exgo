package okex

import (
	"github.com/iaping/exgo/okex/api/account"
)

// https://www.ouyi.fit/docs-v5/zh/#rest-api-account-get-balance
func (c *Client) AccountBalance(ccy string) ([]*account.Balance, error) {
	req := account.NewBalanceRequest(ccy)
	resp := account.NewBalanceResponse()
	if err := c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// https://www.ouyi.fit/docs-v5/zh/#rest-api-account-get-bills-details-last-7-days
func (c *Client) AccountBills(instType, ccy, mgnMode, after, before string, ctType, tp, subType, limit int) ([]*account.Bills, error) {
	req := account.NewBillsRequest(instType, ccy, mgnMode, after, before, ctType, tp, subType, limit)
	resp := account.NewBillsResponse()
	if err := c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// https://www.ouyi.fit/docs-v5/zh/#rest-api-account-get-account-configuration
func (c *Client) AccountConfig() ([]*account.Config, error) {
	req := account.NewConfigRequest()
	resp := account.NewConfigResponse()
	if err := c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// https://www.okx.com/docs-v5/zh/#rest-api-account-get-positions
func (c *Client) AccountPositions(param *account.PositionsParam) ([]*account.Positions, error) {
	req := account.NewPositionsRequest(param)
	resp := account.NewPositionsResponse()

	if err := c.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// https://www.okx.com/docs-v5/zh/#rest-api-account-get-maximum-buy-sell-amount-or-open-amount
func (c *Client) AccountMaxSize(param *account.MaxSizeParam) ([]*account.MaxSize, error) {
	req := account.NewMaxSizeRequest(param)
	resp := account.NewMaxSizeResponse()

	if err := c.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
