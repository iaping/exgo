package market

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const CandlesPath = "/api/v5/market/candles"

type CandlesParam struct {
	InstId string `url:"instId,omitempty"`
	Bar    string `url:"bar,omitempty"`
	After  string `url:"after,omitempty"`
	Before string `url:"before,omitempty"`
	Limit  string `url:"limit,omitempty"`
}

func NewCandlesRequest(param *CandlesParam) *api.CommonRequest {
	return &api.CommonRequest{
		Path:   CandlesPath,
		Method: exgo.HeaderMethodGet,
		Query:  param,
	}
}

type CandlesResponse struct {
	api.CommonResponse
	Data [][]string `json:"data"`
}

func NewCandlesResponse() *CandlesResponse {
	return &CandlesResponse{}
}
