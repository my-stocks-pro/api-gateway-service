package proxy

import "net/http"

type Proxy interface {
	Request(httpMethod string, url string, msg []byte) ([]byte, error)
}

type TypeProxy struct {
	httpClient *http.Client
}

func New(httpClient *http.Client) TypeProxy {
	return TypeProxy{
		httpClient: httpClient,
	}
}
