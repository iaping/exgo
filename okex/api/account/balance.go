package account

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const BalancePath = "/api/v5/account/balance"

type BalanceRequest struct {
	Ccy string `json:"ccy"`
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
	AdjEq   string           `json:"adjEq"`
	Details []*BalanceDetail `json:"details"`
}

type BalanceDetail struct {
	Ccy string `json:"ccy"`
}
