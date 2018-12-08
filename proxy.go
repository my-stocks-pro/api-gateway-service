package main

import (
	"net/http"
	"github.com/kataras/iris/core/errors"
)

type Proxy struct {
	httpClient *http.Client
}

func NewProxy() Proxy {
	return Proxy{
		httpClient: &http.Client{},
	}
}

func (p Proxy) Do(servicePath string) error {
	resp, err := http.Post(servicePath, "application/json", nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}
