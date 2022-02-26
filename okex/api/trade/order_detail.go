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

func NewOrderDetailRequest(param *OrderDetailParam) *api.CommonRequest {
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
	InstType        string `json:"instType"`
	InstId          string `json:"instId"`
	TgtCcy          string `json:"tgtCcy"`
	Ccy             string `json:"ccy"`
	OrdId           string `json:"ordId"`
	ClOrdId         string `json:"clOrdId"`
	Tag             string `json:"tag"`
	Px              string `json:"px"`
	Sz              string `json:"sz"`
	Pnl             string `json:"pnl"`
	OrdType         string `json:"ordType"`
	Side            string `json:"side"`
	PosSide         string `json:"posSide"`
	TdMode          string `json:"tdMode"`
	AccFillSz       string `json:"accFillSz"`
	FillPx          string `json:"fillPx"`
	TradeId         string `json:"tradeId"`
	FillSz          string `json:"fillSz"`
	FillTime        string `json:"fillTime"`
	AvgPx           string `json:"avgPx"`
	State           string `json:"state"`
	Lever           string `json:"lever"`
	TpTriggerPx     string `json:"tpTriggerPx"`
	TpTriggerPxType string `json:"tpTriggerPxType"`
	TpOrdPx         string `json:"tpOrdPx"`
	SlTriggerPx     string `json:"slTriggerPx"`
	SlTriggerPxType string `json:"slTriggerPxType"`
	SlOrdPx         string `json:"slOrdPx"`
	FeeCcy          string `json:"feeCcy"`
	Fee             string `json:"fee"`
	RebateCcy       string `json:"rebateCcy"`
	Source          string `json:"source"`
	Rebate          string `json:"rebate"`
	Category        string `json:"category"`
	UTime           string `json:"uTime"`
	CTime           string `json:"cTime"`
}
