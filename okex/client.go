package okex

import (
	"fmt"

	"github.com/iaping/exgo"
	"github.com/valyala/fasthttp"
)

const (
	Rest             = "https://www.okex.com"
	WebSocketPublic  = "wss://ws.okex.com:8443/ws/v5/public"
	WebSocketPrivate = "wss://ws.okex.com:8443/ws/v5/private"
)

type Client struct {
	Config  *Config
	Request *fasthttp.Client
}

func NewClient(conf *Config) *Client {
	return &Client{
		Config:  conf,
		Request: exgo.FastHttpClient,
	}
}

// request okex api
func (c *Client) Do(api Request) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.Header.SetMethod(api.Method())
	c.setMustHeaders(req)
	if err := c.Request.Do(req, resp); err != nil {
		return nil, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("error http status code: %d", resp.StatusCode())
	}
	return resp.Body(), nil
}

// okex must set headers
func (c *Client) setMustHeaders(req *fasthttp.Request) {
	for k, v := range map[string]string{
		fasthttp.HeaderContentType: exgo.HeaderContentTypeJson,
		fasthttp.HeaderAccept:      exgo.HeaderAcceptJson,
		HeaderAccessKey:            c.Config.APIKey,
		HeaderAccessPassphrase:     c.Config.Passphrase,
		HeaderAccessTimestamp:      "",
		HeaderAccessSign:           "",
	} {
		req.Header.Set(k, v)
	}

	if c.Config.Simulated {
		req.Header.Set(HeaderSimulatedTrading, "1")
	}
}
