package client

import (
	"bytes"
	"net/http"
	"time"
)

var httpClient Doer = &http.Client{
	Timeout: time.Second * 10,
}

var PostClient Poster = &poster{}

func (*poster) Post(url string, header http.Header, body *bytes.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		return nil, err

	}
	req.Header = header

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err

	}
	return res, nil
}
