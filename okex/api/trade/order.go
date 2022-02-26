package trade

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const OrderPath = "/api/v5/trade/order"

type OrderParam struct {
	InstId     string `json:"instId,omitempty"`
	TdMode     string `json:"tdMode,omitempty"`
	Ccy        string `json:"ccy,omitempty"`
	ClOrdId    string `json:"clOrdId,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Side       string `json:"side,omitempty"`
	PosSide    string `json:"posSide,omitempty"`
	OrdType    string `json:"ordType,omitempty"`
	Sz         string `json:"sz,omitempty"`
	Px         string `json:"px,omitempty"`
	ReduceOnly bool   `json:"reduceOnly,omitempty"`
	TgtCcy     string `json:"tgtCcy,omitempty"`
}

func NewOrderRequest(param *OrderParam) *api.CommonRequest {
	return &api.CommonRequest{
		Path:   OrderDetailPath,
		Method: exgo.HeaderMethodPost,
		Body:   param,
	}
}

type OrderResponse struct {
	api.CommonResponse
	Data []*Order `json:"data"`
}

func NewOrderResponse() *OrderResponse {
	return &OrderResponse{}
}

type Order struct {
	OrdId   string `json:"ordId"`
	ClOrdId string `json:"clOrdId"`
	Tag     string `json:"tag"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}
