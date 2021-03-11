package client

import (
	"bytes"
	"net/http"
	"time"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client Doer = &http.Client{
	Timeout: time.Second * 10,
}

func Post(url string, header http.Header, body *bytes.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		return nil, err

	}
	req.Header = header

	res, err := Client.Do(req)

	if err != nil {
		return nil, err

	}
	return res, nil
}
