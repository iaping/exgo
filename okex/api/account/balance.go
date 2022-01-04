package account

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const BalancePath = "/api/v5/account/balance"

type BalanceRequest struct {
	Ccy string `url:"ccy"`
}

func NewBalanceRequest(ccy string) *BalanceRequest {
	return &BalanceRequest{Ccy: ccy}
}

func (br *BalanceRequest) Path() string {
	return BalancePath
}

func (br *BalanceRequest) Method() string {
	return exgo.HeaderMethodGet
}

type BalanceResponse struct {
	api.CommonResponse
	Data []*Balance `json:"data"`
}

type Balance struct {
	AdjEq       string           `json:"adjEq,omitempty"`
	Imr         string           `json:"imr,omitempty"`
	IsoEq       string           `json:"isoEq,omitempty"`
	MgnRatio    string           `json:"mgnRatio,omitempty"`
	Mmr         string           `json:"mmr,omitempty"`
	NotionalUsd string           `json:"notionalUsd,omitempty"`
	OrdFroz     string           `json:"ordFroz,omitempty"`
	TotalEq     string           `json:"totalEq"`
	UTime       string           `json:"uTime"`
	Details     []*BalanceDetail `json:"details"`
}

type BalanceDetail struct {
	AvailBal      string `json:"availBal,omitempty"`
	AvailEq       string `json:"availEq,omitempty"`
	CashBal       string `json:"cashBal"`
	Ccy           string `json:"ccy"`
	CrossLiab     string `json:"crossLiab,omitempty"`
	DisEq         string `json:"disEq"`
	Eq            string `json:"eq"`
	EqUsd         string `json:"eqUsd"`
	FrozenBal     string `json:"FrozenBal"`
	Interest      string `json:"interest,omitempty"`
	IsoEq         string `json:"isoEq,omitempty"`
	IsoLiab       string `json:"isoLiab,omitempty"`
	IsoUpl        string `json:"isoUpl,omitempty"`
	Liab          string `json:"liab,omitempty"`
	MaxLoan       string `json:"maxLoan,omitempty"`
	MgnRatio      string `json:"mgnRatio,omitempty"`
	NotionalLever string `json:"notionalLever,omitempty"`
	OrdFrozen     string `json:"ordFrozen"`
	StgyEq        string `json:"stgyEq,omitempty"`
	Twap          string `json:"twap,omitempty"`
	UTime         string `json:"uTime"`
	Upl           string `json:"upl,omitempty"`
	UplLiab       string `json:"uplLiab,omitempty"`
}
