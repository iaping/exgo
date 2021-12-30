package exgo

import (
	"time"

	"github.com/valyala/fasthttp"
)

const (
	HeaderContentTypeJson = "application/json;charset=utf-8"
	HeaderAcceptJson      = "application/json"
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
