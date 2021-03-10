package rest

import (
	"bytes"
	"net/http"
)

func Post(url string, header http.Header, body *bytes.Reader) (*http.Response, error) {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		return nil, err

	}
	req.Header = header

	res, err := client.Do(req)

	if err != nil {
		return nil, err

	}
	return res, nil
}
