package infrastructure

import (
	"net/http"
	"time"
)

func GetHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(5) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}