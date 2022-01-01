package okex

import (
	"encoding/json"
	"fmt"

	"github.com/iaping/exgo"
	"github.com/iaping/exgo/okex/api"
	"github.com/iaping/exgo/okex/api/account"
	"github.com/iaping/exgo/okex/api/asset"
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
	data, err := c.do(req)
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
func (c *Client) do(r api.Request) ([]byte, error) {
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
	sign := c.signature(r)
	signature, err := sign.String()
	if err != nil {
		return err
	}
	necessaries := map[string]string{
		fasthttp.HeaderContentType: exgo.HeaderContentTypeJson,
		fasthttp.HeaderAccept:      exgo.HeaderAcceptJson,
		HeaderAccessKey:            c.Config.APIKey,
		HeaderAccessPassphrase:     c.Config.Passphrase,
		HeaderAccessTimestamp:      sign.Timestamp,
		HeaderAccessSign:           signature,
	}
	if c.Config.Simulated {
		necessaries[HeaderSimulatedTrading] = "1"
	}
	for k, v := range necessaries {
		req.Header.Set(k, v)
	}
	req.Header.SetMethod(r.Method())
	req.SetRequestURI(c.Host + r.Path())
	req.SetBodyString(sign.Body)
	return nil
}

// signature Request
func (c *Client) signature(r api.Request) *api.Signature {
	body, _ := json.Marshal(r)
	return api.NewSignature(c.Config.SecretKey, r.Method(), r.Path(), string(body))
}

func (c *Client) Currencies() ([]*asset.Currencies, error) {
	req := asset.NewCurrenciesRequest()
	var resp asset.CurrenciesResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) Balance(ccy string) ([]*account.Balance, error) {
	req := account.NewBalanceRequest(ccy)
	var resp account.BalanceResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
