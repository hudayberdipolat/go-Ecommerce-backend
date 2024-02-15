package http

import (
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     30 * time.Second,
			MaxIdleConns:        30,
			TLSHandshakeTimeout: time.Minute * 1,
		},
		Timeout: time.Second * 30,
	}
}
