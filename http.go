package exgo

import (
	"time"

	"github.com/valyala/fasthttp"
)

var (
	FastHttpClient = &fasthttp.Client{
		Name:                "exgo",
		MaxConnsPerHost:     16,
		MaxIdleConnDuration: 20 * time.Second,
		ReadTimeout:         10 * time.Second,
		WriteTimeout:        10 * time.Second,
	}
)
