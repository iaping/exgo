package okex

import (
	"github.com/iaping/exgo/okex/api/trade"
)

// https://www.okx.com/docs-v5/zh/#rest-api-trade-place-order
func (c *Client) TradeOrder(param *trade.OrderParam) ([]*trade.Order, error) {
	req := trade.NewOrderRequest(param)
	resp := trade.NewOrderResponse()

	if err := c.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// https://www.okx.com/docs-v5/zh/#rest-api-trade-get-order-details
func (c *Client) TradeOrderDetail(param *trade.OrderDetailParam) ([]*trade.OrderDetail, error) {
	req := trade.NewOrderDetailRequest(param)
	resp := trade.NewOrderDetailResponse()

	if err := c.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
