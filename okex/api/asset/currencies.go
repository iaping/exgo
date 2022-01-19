package asset

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const CurrenciesPath = "/api/v5/asset/currencies"

func NewCurrenciesRequest() *api.CommonRequest {
	return &api.CommonRequest{
		Path:   CurrenciesPath,
		Method: exgo.HeaderMethodGet,
	}
}

type CurrenciesResponse struct {
	api.CommonResponse
	Data []*Currencies `json:"data"`
}

func NewCurrenciesResponse() *CurrenciesResponse {
	return &CurrenciesResponse{}
}

type Currencies struct {
	Ccy         string  `json:"ccy"`
	Chain       string  `json:"chain"`
	Name        string  `json:"name"`
	CanDep      bool    `json:"canDep"`
	CanWd       bool    `json:"canWd"`
	CanInternal bool    `json:"canInternal"`
	MinWd       float64 `json:"minWd,string"`
	MaxFee      float64 `json:"maxFee,string"`
	MinFee      float64 `json:"minFee,string"`
	MainNet     bool    `json:"mainNet"`
}
