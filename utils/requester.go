package utils

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
)

func NewRequest(url string, b []byte) ([]byte, error) {

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}


	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, err
	}

	body, errReadBody := ioutil.ReadAll(resp.Body)
	if errReadBody != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return body, nil
}

