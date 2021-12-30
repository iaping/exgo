package okex

type Response interface {
	Code() string
	Msg() string
	Data() interface{}
}
