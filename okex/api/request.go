package api

// okex api request
type Request interface {
	GetPath() string
	GetMethod() string
	GetQuery() interface{}
	GetBody() interface{}
}

type CommonRequest struct {
	Path   string      `json:"path"`
	Method string      `json:"method"`
	Query  interface{} `json:"query"`
	Body   interface{} `json:"body"`
}

func (r *CommonRequest) GetPath() string {
	return r.Path
}

func (r *CommonRequest) GetMethod() string {
	return r.Method
}

func (r *CommonRequest) GetQuery() interface{} {
	return r.Query
}

func (r *CommonRequest) GetBody() interface{} {
	return r.Body
}
