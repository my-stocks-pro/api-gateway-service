package proxy

import (
	"net/http"
	"io/ioutil"
	"github.com/kataras/iris/core/errors"
	"io"
)

func (p TypeProxy) Request(httpMethod string, url string, body io.ReadCloser) (io.ReadCloser, error) {
	req, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return resp.Body, nil
}

func (p TypeProxy) ReadResponseBody(blob io.ReadCloser) ([]byte, error) {
	body, err := ioutil.ReadAll(blob)
	if err != nil {
		return nil, err
	}

	return body, nil
}
