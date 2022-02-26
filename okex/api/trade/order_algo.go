package trade

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const OrderAlgoPath = "/api/v5/trade/order-algo"

type OrderAlgoParam struct {
	InstId     string `json:"instId,omitempty"`
	TdMode     string `json:"tdMode,omitempty"`
	Ccy        string `json:"ccy,omitempty"`
	Side       string `json:"side,omitempty"`
	PosSide    string `json:"posSide,omitempty"`
	OrdType    string `json:"ordType,omitempty"`
	Sz         string `json:"sz,omitempty"`
	Tag        string `json:"tag,omitempty"`
	TgtCcy     string `json:"tgtCcy,omitempty"`
	ReduceOnly bool   `json:"reduceOnly,omitempty"`

	// 止盈止损
	TpTriggerPx     string `json:"tpTriggerPx,omitempty"`
	TpTriggerPxType string `json:"tpTriggerPxType,omitempty"`
	TpOrdPx         string `json:"tpOrdPx,omitempty"`
	SlTriggerPx     string `json:"slTriggerPx,omitempty"`
	SlTriggerPxType string `json:"slTriggerPxType,omitempty"`
	SlOrdPx         string `json:"slOrdPx,omitempty"`
}

func NewOrderAlgoRequest(param *OrderAlgoParam) *api.CommonRequest {
	return &api.CommonRequest{
		Path:   OrderAlgoPath,
		Method: exgo.HeaderMethodPost,
		Body:   param,
	}
}

type OrderAlgoResponse struct {
	api.CommonResponse
	Data []*OrderAlgo `json:"data"`
}

func NewOrderAlgoResponse() *OrderAlgoResponse {
	return &OrderAlgoResponse{}
}

type OrderAlgo struct {
	AlgoId string `json:"algoId"`
	SCode  string `json:"sCode"`
	SMsg   string `json:"sMsg"`
}
