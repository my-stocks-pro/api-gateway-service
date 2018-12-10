package proxy

import "net/http"

type Proxy interface {
	POST(servicePath string, body []byte) error
}

type TypeProxy struct {
	httpClient *http.Client
}

func New(httpClient *http.Client) TypeProxy {
	return TypeProxy{
		httpClient: httpClient,
	}
}
