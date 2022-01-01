package api

// okex api request
type Request interface {
	Path() string
	Method() string
}
