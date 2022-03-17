package public

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const instrumentsPath = "/api/v5/public/instruments"

type InstrumentsParam struct {
	InstType string `url:"instType"`
	Uly      string `url:"uly,omitempty"`
	InstId   string `url:"instId,omitempty"`
}

func NewInstrumentsRequest(param *InstrumentsParam) *api.CommonRequest {
	return &api.CommonRequest{
		Path:   instrumentsPath,
		Method: exgo.HeaderMethodGet,
		Query:  param,
	}
}

type InstrumentsResponse struct {
	api.CommonResponse
	Data []*Instruments `json:"data"`
}

func NewInstrumentsResponse() *InstrumentsResponse {
	return &InstrumentsResponse{}
}

type Instruments struct {
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	Uly       string `json:"uly"`
	Category  string `json:"category"`
	BaseCcy   string `json:"baseCcy"`
	QuoteCcy  string `json:"quoteCcy"`
	SettleCcy string `json:"settleCcy"`
	CtVal     string `json:"ctVal"`
	CtMult    string `json:"ctMult"`
	CtValCcy  string `json:"ctValCcy"`
	OptType   string `json:"optType"`
	Stk       string `json:"stk"`
	ListTime  string `json:"listTime"`
	ExpTime   string `json:"expTime"`
	Lever     string `json:"lever"`
	TickSz    string `json:"tickSz"`
	LotSz     string `json:"lotSz"`
	MinSz     string `json:"minSz"`
	CtType    string `json:"ctType"`
	Alias     string `json:"alias"`
	State     string `json:"state"`
}
