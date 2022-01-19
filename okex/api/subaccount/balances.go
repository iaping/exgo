package subaccount

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const SubaccountBalancesPath = "/api/v5/account/subaccount/balances"

type SubaccountBalancesParam struct {
	SubAcct string `url:"subAcct"`
}

func NewSubaccountBalancesRequest(subAcct string) *api.CommonRequest {
	param := &SubaccountBalancesParam{
		SubAcct: subAcct,
	}

	return &api.CommonRequest{
		Path:   SubaccountBalancesPath,
		Method: exgo.HeaderMethodGet,
		Query:  param,
	}
}
