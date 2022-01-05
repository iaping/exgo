package subaccount

import (
	"github.com/iaping/exgo"
)

const SubaccountBalancesPath = "/api/v5/account/subaccount/balances"

type SubaccountBalancesRequest struct {
	SubAcct string `url:"subAcct"`
}

func NewSubaccountBalancesRequest(subAcct string) *SubaccountBalancesRequest {
	return &SubaccountBalancesRequest{SubAcct: subAcct}
}

func (r *SubaccountBalancesRequest) Path() string {
	return SubaccountBalancesPath
}

func (r *SubaccountBalancesRequest) Method() string {
	return exgo.HeaderMethodGet
}
