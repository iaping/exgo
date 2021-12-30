package asset

import "github.com/iaping/exgo/okex"

const BalancesRequestPath = "/api/v5/asset/balances"

type BalancesRequest struct {
	Ccy string
}

func (br *BalancesRequest) RequestPath() string {
	return BalancesRequestPath
}

func (br *BalancesRequest) Method() string {
	return okex.HeaderMethodGet
}

type BalancesResponse struct {
	Ccy       string
	Bal       string
	FrozenBal string
	AvailBal  string
}
