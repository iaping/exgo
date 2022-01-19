package account

import (
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
)

const ConfigPath = "/api/v5/account/config"

func NewConfigRequest() *api.CommonRequest {
	return &api.CommonRequest{
		Path:   ConfigPath,
		Method: exgo.HeaderMethodGet,
	}
}

type ConfigResponse struct {
	api.CommonResponse
	Data []*Config `json:"data"`
}

func NewConfigResponse() *ConfigResponse {
	return &ConfigResponse{}
}

type Config struct {
	Uid        uint64 `json:"uid,string"`
	AcctLv     int    `json:"acctLv,string"`
	PosMode    string `json:"posMode"`
	AutoLoan   bool   `json:"autoLoan"`
	GreeksType string `json:"greeksType"`
	Level      string `json:"level"`
	LevelTmp   string `json:"levelTmp"`
	CtIsoMode  string `json:"ctIsoMode"`
	MgnIsoMode string `json:"mgnIsoMode"`
}
