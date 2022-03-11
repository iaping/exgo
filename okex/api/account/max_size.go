package account

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const MaxSizePath = "/api/v5/account/max-size"

type MaxSizeParam struct {
	InstId   string `url:"instId,omitempty"`
	TdMode   string `url:"tdMode,omitempty"`
	Ccy      string `url:"ccy,omitempty"`
	Px       string `url:"px,omitempty"`
	Leverage string `url:"leverage,omitempty"`
}

func NewMaxSizeRequest(param *MaxSizeParam) *api.CommonRequest {
	return &api.CommonRequest{
		Path:   MaxSizePath,
		Method: exgo.HeaderMethodGet,
		Query:  param,
	}
}

type MaxSizeResponse struct {
	api.CommonResponse
	Data []*MaxSize `json:"data"`
}

func NewMaxSizeResponse() *MaxSizeResponse {
	return &MaxSizeResponse{}
}

type MaxSize struct {
	InstId  string `json:"instId"`
	Ccy     string `json:"ccy"`
	MaxBuy  string `json:"maxBuy"`
	MaxSell string `json:"maxSell"`
}
