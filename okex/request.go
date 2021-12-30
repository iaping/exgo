package okex

type Request interface {
	RequestPath() string
	Method() string
}
