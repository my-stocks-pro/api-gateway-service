package proxy

import (
	"net/http"
	"io"
)

type Proxy interface {
	Request(httpMethod string, url string, body io.ReadCloser) (io.ReadCloser, error)
	ReadResponseBody(blob io.ReadCloser) ([]byte, error)
}

type TypeProxy struct {
	httpClient *http.Client
}

func New(httpClient *http.Client) TypeProxy {
	return TypeProxy{
		httpClient: httpClient,
	}
}
