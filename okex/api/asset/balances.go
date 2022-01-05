package asset

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const BalancesPath = "/api/v5/asset/balances"

type BalancesRequest struct {
	Ccy string `url:"ccy,omitempty"`
}

func NewBalancesRequest(ccy string) *BalancesRequest {
	return &BalancesRequest{
		Ccy: ccy,
	}
}

func (r *BalancesRequest) Path() string {
	return BalancesPath
}

func (r *BalancesRequest) Method() string {
	return exgo.HeaderMethodGet
}

type BalancesResponse struct {
	api.CommonResponse
	Data []*Balances `json:"data"`
}

type Balances struct {
	Ccy       string  `json:"ccy"`
	Bal       float64 `json:"bal,string"`
	FrozenBal float64 `json:"frozenBal,string"`
	AvailBal  float64 `json:"availBal,string"`
}
