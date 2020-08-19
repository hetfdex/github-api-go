package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Post request to a given url
func Post(url string, body interface{}, header http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = header

	client := http.Client{}

	return client.Do(request)
}
