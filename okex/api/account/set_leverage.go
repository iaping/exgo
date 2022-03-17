package account

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const setLeveragePath = "/api/v5/account/set-leverage"

type SetLeverageParam struct {
	InstId  string `json:"instId,omitempty"`
	Ccy     string `json:"ccy,omitempty"`
	Lever   string `json:"lever,omitempty"`
	MgnMode string `json:"mgnMode,omitempty"`
	PosSide string `json:"posSide,omitempty"`
}

func NewSetLeverageRequest(param *SetLeverageParam) *api.CommonRequest {
	return &api.CommonRequest{
		Path:   setLeveragePath,
		Method: exgo.HeaderMethodPost,
		Body:   param,
	}
}

type SetLeverageResponse struct {
	api.CommonResponse
	Data []*SetLeverage `json:"data"`
}

func NewSetLeverageResponse() *SetLeverageResponse {
	return &SetLeverageResponse{}
}

type SetLeverage struct {
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
	InstId  string `json:"instId"`
	PosSide string `json:"posSide"`
}
