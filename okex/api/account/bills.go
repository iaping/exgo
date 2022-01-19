package account

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const BillsPath = "/api/v5/account/bills"

type BillsParam struct {
	InstType string `url:"instType,omitempty"`
	Ccy      string `url:"ccy,omitempty"`
	MgnMode  string `url:"mgnMode,omitempty"`
	CtType   int    `url:"ctType,omitempty"`
	Type     int    `url:"type,omitempty"`
	SubType  int    `url:"subType,omitempty"`
	After    string `url:"after,omitempty"`
	Before   string `url:"before,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

func NewBillsRequest(instType, ccy, mgnMode, after, before string, ctType, tp, subType, limit int) *api.CommonRequest {
	param := &BillsParam{
		InstType: instType,
		Ccy:      ccy,
		MgnMode:  mgnMode,
		CtType:   ctType,
		Type:     tp,
		SubType:  subType,
		After:    after,
		Before:   before,
		Limit:    limit,
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
	InstType  string `json:"instType"`
	BillId    string `json:"billId"`
	Type      int    `json:"type,string"`
	SubType   int    `json:"subType,string"`
	Ts        uint   `json:"ts,string"`
	BalChg    string `json:"balChg"`
	PosBalChg string `json:"posBalChg"`
	Bal       string `json:"bal"`
	PosBal    string `json:"posBal"`
	Sz        string `json:"sz"`
	Ccy       string `json:"ccy"`
	Pnl       string `json:"pnl"`
	Fee       string `json:"fee"`
	MgnMode   string `json:"mgnMode"`
	InstId    string `json:"instId"`
	OrdId     string `json:"ordId"`
	ExecType  string `json:"execType"`
	From      string `json:"from"`
	To        string `json:"to"`
	Notes     string `json:"notes"`
}
