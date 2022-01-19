package asset

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const BalancesPath = "/api/v5/asset/balances"

type BalancesParam struct {
	Ccy string `url:"ccy,omitempty"`
}

func NewBalancesRequest(ccy string) *api.CommonRequest {
	param := &BalancesParam{
		Ccy: ccy,
	}

	return &api.CommonRequest{
		Path:   BalancesPath,
		Method: exgo.HeaderMethodGet,
		Query:  param,
	}
}

type BalancesResponse struct {
	api.CommonResponse
	Data []*Balances `json:"data"`
}

func NewBalancesResponse() *BalancesResponse {
	return &BalancesResponse{}
}

type Balances struct {
	Ccy       string  `json:"ccy"`
	Bal       float64 `json:"bal,string"`
	FrozenBal float64 `json:"frozenBal,string"`
	AvailBal  float64 `json:"availBal,string"`
}
