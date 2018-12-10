package proxy

import (
	"net/http"
	"github.com/kataras/iris/core/errors"
	"bytes"
)

func (p TypeProxy) POST(servicePath string, body []byte) error {
	resp, err := http.Post(servicePath, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}
