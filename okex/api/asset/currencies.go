package asset

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const CurrenciesPath = "/api/v5/asset/currencies"

type CurrenciesRequest struct{}

func NewCurrenciesRequest() *CurrenciesRequest {
	return &CurrenciesRequest{}
}

func (br *CurrenciesRequest) Path() string {
	return CurrenciesPath
}

func (br *CurrenciesRequest) Method() string {
	return exgo.HeaderMethodGet
}

type CurrenciesResponse struct {
	api.CommonResponse
	Data []*Currencies `json:"data"`
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
