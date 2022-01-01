package api

type Response interface {
	IsSuccess() bool
	GetCode() string
	GetMessage() string
}

type CommonResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (resp CommonResponse) IsSuccess() bool {
	return resp.Code == "0"
}

func (resp CommonResponse) GetCode() string {
	return resp.Code
}

func (resp CommonResponse) GetMessage() string {
	return resp.Message
}
