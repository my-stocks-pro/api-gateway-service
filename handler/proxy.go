package handler

import (
	"net/http"
	"github.com/kataras/iris/core/errors"
	"bytes"
)

type Proxy struct {
	httpClient *http.Client
}

func NewProxy() Proxy {
	return Proxy{
		httpClient: &http.Client{},
	}
}

func (p Proxy) POST(servicePath string, body []byte) error {
	resp, err := http.Post(servicePath, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}
