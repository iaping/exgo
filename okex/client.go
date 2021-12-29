package okex

import (
	"github.com/iaping/exgo"
	"github.com/valyala/fasthttp"
)

const (
	Rest             = "https://www.okex.com"
	WebSocketPublic  = "wss://ws.okex.com:8443/ws/v5/public"
	WebSocketPrivate = "wss://ws.okex.com:8443/ws/v5/private"
)

type Client struct {
	Config *Config
	Http   *fasthttp.Client
}

func NewClient(conf *Config) *Client {
	return &Client{
		Config: conf,
		Http:   exgo.FastHttpClient,
	}
}

func (c *Client) Do() {

}
