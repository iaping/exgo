package account

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const BalancePath = "/api/v5/account/balance"

type BalanceParam struct {
	Ccy string `url:"ccy"`
}

func NewBalanceRequest(ccy string) *api.CommonRequest {
	param := &BalanceParam{
		Ccy: ccy,
	}

	return &api.CommonRequest{
		Path:   BalancePath,
		Method: exgo.HeaderMethodGet,
		Query:  param,
	}
}

type BalanceResponse struct {
	api.CommonResponse
	Data []*Balance `json:"data"`
}

func NewBalanceResponse() *BalanceResponse {
	return &BalanceResponse{}
}

type Balance struct {
	AdjEq       string           `json:"adjEq,omitempty"`
	Imr         string           `json:"imr,omitempty"`
	IsoEq       string           `json:"isoEq,omitempty"`
	MgnRatio    string           `json:"mgnRatio,omitempty"`
	Mmr         string           `json:"mmr,omitempty"`
	NotionalUsd string           `json:"notionalUsd,omitempty"`
	OrdFroz     string           `json:"ordFroz,omitempty"`
	TotalEq     float64          `json:"totalEq,string"`
	UTime       uint             `json:"uTime,string"`
	Details     []*BalanceDetail `json:"details"`
}

type BalanceDetail struct {
	AvailBal      string  `json:"availBal,omitempty"`
	AvailEq       string  `json:"availEq,omitempty"`
	CashBal       float64 `json:"cashBal,string"`
	Ccy           string  `json:"ccy"`
	CrossLiab     string  `json:"crossLiab,omitempty"`
	DisEq         float64 `json:"disEq,string"`
	Eq            float64 `json:"eq,string"`
	EqUsd         float32 `json:"eqUsd,string"`
	FrozenBal     float64 `json:"FrozenBal,string"`
	Interest      string  `json:"interest,omitempty"`
	IsoEq         string  `json:"isoEq,omitempty"`
	IsoLiab       string  `json:"isoLiab,omitempty"`
	IsoUpl        string  `json:"isoUpl,omitempty"`
	Liab          string  `json:"liab,omitempty"`
	MaxLoan       string  `json:"maxLoan,omitempty"`
	MgnRatio      string  `json:"mgnRatio,omitempty"`
	NotionalLever string  `json:"notionalLever,omitempty"`
	OrdFrozen     float64 `json:"ordFrozen,string"`
	StgyEq        string  `json:"stgyEq,omitempty"`
	Twap          string  `json:"twap,omitempty"`
	UTime         uint    `json:"uTime,string"`
	Upl           string  `json:"upl,omitempty"`
	UplLiab       string  `json:"uplLiab,omitempty"`
}
