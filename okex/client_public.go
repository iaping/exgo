package okex

import (
	"github.com/iaping/exgo/okex/api/public"
)

// https://www.okx.com/docs-v5/zh/#rest-api-public-data-get-instruments
func (c *Client) PublicInstruments(param *public.InstrumentsParam) ([]*public.Instruments, error) {
	req := public.NewInstrumentsRequest(param)
	resp := public.NewInstrumentsResponse()

	if err := c.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
