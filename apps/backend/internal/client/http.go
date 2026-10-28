package client

import (
	"net/http"
	"time"
)

type HttpAPI interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpAPI struct {
	client *http.Client
}

func NewHTTPAPI(timeout time.Duration) HttpAPI {
	return &httpAPI{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (h *httpAPI) Do(req *http.Request) (*http.Response, error) {
	// #nosec G704 -- callers pass fixed, hardcoded API URLs; no user-controlled host
	return h.client.Do(req)
}
