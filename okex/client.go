package okex

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
	"github.com/valyala/fasthttp"
)

const (
	Host      = "https://www.okex.com"
	HostChina = "https://www.ouyi.fit"
)

type Client struct {
	Host    string
	Config  *Config
	Request *fasthttp.Client
}

// new okex client
func New(apiKey, secretKey, passphrase string, simulated, china bool) *Client {
	cnf := &Config{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		Passphrase: passphrase,
		Simulated:  simulated,
		China:      china,
	}

	return NewClient(cnf)
}

// new okex client
func NewClient(cnf *Config) *Client {
	client := &Client{
		Host:    Host,
		Config:  cnf,
		Request: exgo.FastHttpClient,
	}

	if cnf.China {
		client.Host = HostChina
	}

	return client
}

// request okex api and handle the response
func (c *Client) Do(req api.Request, resp api.Response) error {
	data, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, resp); err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("response code: %s, message: %s", resp.GetCode(), resp.GetMessage())
	}

	return nil
}

// request okex api
func (c *Client) doRequest(api api.Request) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	c.handleRequest(req, c.signature(api))

	if err := c.Request.Do(req, resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("http status code:%d, desc:%s", resp.StatusCode(), string(resp.Body()))
	}

	return resp.Body(), nil
}

// okex must set header
func (c *Client) handleRequest(req *fasthttp.Request, signature *signature) {
	headers := map[string]string{
		fasthttp.HeaderContentType: exgo.HeaderContentTypeJson,
		fasthttp.HeaderAccept:      exgo.HeaderAcceptJson,
		HeaderAccessKey:            c.Config.APIKey,
		HeaderAccessPassphrase:     c.Config.Passphrase,
		HeaderAccessSign:           signature.Build(),
		HeaderAccessTimestamp:      signature.Timestamp,
	}
	if c.Config.Simulated {
		headers[HeaderSimulatedTrading] = "1"
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Header.SetMethod(signature.Method)
	req.SetRequestURI(c.Host + signature.Path)

	if signature.Body != "" {
		req.SetBodyString(signature.Body)
	}
}

// signature Request
func (c *Client) signature(api api.Request) *signature {
	body := []byte{}
	if api.GetBody() != nil {
		body, _ = json.Marshal(api.GetBody())
	}

	path := api.GetPath()
	if api.GetQuery() != nil {
		values, _ := query.Values(api.GetQuery())
		if len(values) > 0 {
			path += "?" + values.Encode()
		}
	}

	return signing(c.Config.SecretKey, api.GetMethod(), path, string(body))
}
