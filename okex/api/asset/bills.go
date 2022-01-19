package asset

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const BillsPath = "/api/v5/asset/bills"

type BillsParam struct {
	Ccy    string `url:"ccy,omitempty"`
	Type   int    `url:"type,omitempty"`
	After  int64  `url:"after,omitempty"`
	Before int64  `url:"before,omitempty"`
	Limit  int    `url:"limit,omitempty"`
}

func NewBillsRequest(ccy string, typed, limit int, after, before int64) *api.CommonRequest {
	param := &BillsParam{
		Ccy:    ccy,
		Type:   typed,
		After:  after,
		Before: before,
		Limit:  limit,
	}

	return &api.CommonRequest{
		Path:   BillsPath,
		Method: exgo.HeaderMethodGet,
		Query:  param,
	}
}

type BillsResponse struct {
	api.CommonResponse
	Data []*Bills `json:"data"`
}

func NewBillsResponse() *BillsResponse {
	return &BillsResponse{}
}

type Bills struct {
	Ccy    string  `json:"ccy"`
	BillId string  `json:"billId"`
	BalChg float64 `json:"balChg,string"`
	Bal    float64 `json:"bal,string"`
	Type   int     `json:"type,string"`
	Ts     int64   `json:"ts,string"`
}
