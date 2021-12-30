package okex

type Request interface {
	Path() string
	Method() string
}
