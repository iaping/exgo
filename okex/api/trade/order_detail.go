package trade

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const OrderDetailPath = "/api/v5/trade/order"

type OrderDetailParam struct {
	InstId  string `url:"instId,omitempty"`
	OrdId   string `url:"ordId,omitempty"`
	ClOrdId string `url:"clOrdId,omitempty"`
}

func NewOrderDetailRequest(instId, ordId, clOrdId string) *api.CommonRequest {
	param := &OrderDetailParam{
		InstId:  instId,
		OrdId:   ordId,
		ClOrdId: clOrdId,
	}

	return &api.CommonRequest{
		Path:   OrderDetailPath,
		Method: exgo.HeaderMethodGet,
		Query:  param,
	}
}

type OrderDetailResponse struct {
	api.CommonResponse
	Data []*OrderDetail `json:"data"`
}

func NewOrderDetailResponse() *OrderDetailResponse {
	return &OrderDetailResponse{}
}

type OrderDetail struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	TgtCcy   string `json:"tgtCcy"`
	Ccy      string `json:"ccy"`
	OrdId    string `json:"ordId"`
	ClOrdId  string `json:"clOrdId"`
}
