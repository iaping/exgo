package account

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const PositionsPath = "/api/v5/account/positions"

type PositionsRequest struct{}

func NewPositionsRequest(ccy string) *PositionsRequest {
	return &PositionsRequest{}
}

func (br *PositionsRequest) Path() string {
	return PositionsPath
}

func (br *PositionsRequest) Method() string {
	return exgo.HeaderMethodGet
}

type PositionsResponse struct {
	api.CommonResponse
	Data []*Positions `json:"data"`
}

type Positions struct {
	AdjEq string `json:"adjEq"`
	Imr   string `json:"imr"`
}
