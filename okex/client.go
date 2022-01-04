package okex

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
	"github.com/valyala/fasthttp"
)

const (
	Host = "https://www.ouyi.fit"
)

type Client struct {
	Host    string
	Config  *Config
	Request *fasthttp.Client
}

func NewClient(conf *Config) *Client {
	return &Client{
		Host:    Host,
		Config:  conf,
		Request: exgo.FastHttpClient,
	}
}

// request okex api
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
func (c *Client) doRequest(r api.Request) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	if err := c.setNecessary(req, r); err != nil {
		return nil, err
	}
	if err := c.Request.Do(req, resp); err != nil {
		return nil, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("http status code:%d, desc:%s", resp.StatusCode(), string(resp.Body()))
	}
	return resp.Body(), nil
}

// okex necessary set
func (c *Client) setNecessary(req *fasthttp.Request, r api.Request) error {
	signature := c.signature(r)
	necessaries := map[string]string{
		fasthttp.HeaderContentType: exgo.HeaderContentTypeJson,
		fasthttp.HeaderAccept:      exgo.HeaderAcceptJson,
		HeaderAccessKey:            c.Config.APIKey,
		HeaderAccessPassphrase:     c.Config.Passphrase,
		HeaderAccessSign:           signature.Build(),
		HeaderAccessTimestamp:      signature.Timestamp,
	}
	if c.Config.Simulated {
		necessaries[HeaderSimulatedTrading] = "1"
	}
	for k, v := range necessaries {
		req.Header.Set(k, v)
	}
	req.Header.SetMethod(r.Method())
	req.SetBodyString(signature.Body)
	req.SetRequestURI(c.Host + signature.Path)
	return nil
}

// signature Request
func (c *Client) signature(r api.Request) *api.Signature {
	path := r.Path()
	body := []byte("")
	if c.isPost(r) {
		body, _ = c.buildBody(r)
	} else {
		q, _ := c.buildQuery(r)
		path += "?" + q.Encode()
	}
	return api.NewSignature(c.Config.SecretKey, r.Method(), path, string(body))
}

// build request body
func (c *Client) buildBody(r api.Request) ([]byte, error) {
	return json.Marshal(r)
}

// build request uri query
func (c *Client) buildQuery(r api.Request) (url.Values, error) {
	return query.Values(r)
}

// is post request
func (c *Client) isPost(r api.Request) bool {
	return r.Method() == "POST"
}
