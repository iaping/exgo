package account

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const PositionsPath = "/api/v5/account/positions"

type PositionsParam struct {
	InstId   string `url:"instId,omitempty"`
	InstType string `url:"instType,omitempty"`
	PosId    string `url:"posId,omitempty"`
}

func NewPositionsRequest(param *PositionsParam) *api.CommonRequest {
	return &api.CommonRequest{
		Path:   PositionsPath,
		Method: exgo.HeaderMethodGet,
		Query:  param,
	}
}

type PositionsResponse struct {
	api.CommonResponse
	Data []*Positions `json:"data"`
}

func NewPositionsResponse() *PositionsResponse {
	return &PositionsResponse{}
}

type Positions struct {
	InstType    string `json:"instType"`
	MgnMode     string `json:"mgnMode"`
	PosId       string `json:"posId"`
	PosSide     string `json:"posSide"`
	Pos         string `json:"pos"`
	BaseBal     string `json:"baseBal"`
	QuoteBal    string `json:"quoteBal"`
	PosCcy      string `json:"posCcy"`
	AvailPos    string `json:"availPos"`
	AvgPx       string `json:"avgPx"`
	Upl         string `json:"upl"`
	UplRatio    string `json:"uplRatio"`
	InstId      string `json:"instId"`
	Lever       string `json:"lever"`
	LiqPx       string `json:"liqPx"`
	MarkPx      string `json:"markPx"`
	Imr         string `json:"imr"`
	Margin      string `json:"margin"`
	MgnRatio    string `json:"mgnRatio"`
	Mmr         string `json:"mmr"`
	Liab        string `json:"liab"`
	LiabCcy     string `json:"liabCcy"`
	Interest    string `json:"interest"`
	TradeId     string `json:"tradeId"`
	OptVal      string `json:"optVal"`
	NotionalUsd string `json:"notionalUsd"`
	Adl         string `json:"adl"`
	Ccy         string `json:"ccy"`
	Last        string `json:"last"`
	UsdPx       string `json:"usdPx"`
	DeltaBS     string `json:"deltaBS"`
	DeltaPA     string `json:"deltaPA"`
	GammaBS     string `json:"gammaBS"`
	GammaPA     string `json:"gammaPA"`
	ThetaBS     string `json:"thetaBS"`
	ThetaPA     string `json:"thetaPA"`
	VegaBS      string `json:"vegaBS"`
	VegaPA      string `json:"vegaPA"`
	CTime       string `json:"cTime"`
	UTime       string `json:"uTime"`
}
