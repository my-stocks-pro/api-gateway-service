package proxy

import (
	"net/http"
	"bytes"
	"io/ioutil"
)

func (p TypeProxy) Request(httpMethod string, url string, msg []byte) ([]byte, error) {
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(msg))
	if err != nil {
		return nil, err
	}

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
