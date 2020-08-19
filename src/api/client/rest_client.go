package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var mocking = false

var mocks = make(map[string]*Mock)

// Mock model
type Mock struct {
	URL        string
	HTTPMethod string
	Response   *http.Response
	Err        error
}

// StartMocker enables client mocking
func StartMocker() {
	mocking = true
}

// StopMocker disables client mocking
func StopMocker() {
	mocking = false
}

func getMockID(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

// AddMock inserts mock with url key
func AddMock(mock Mock) {
	mocks[getMockID(mock.HTTPMethod, mock.URL)] = &mock
}

// Post request to a given url
func Post(url string, body interface{}, header http.Header) (*http.Response, error) {
	if mocking {
		mock := mocks[getMockID(http.MethodPost, url)]

		if mock == nil {
			return nil, fmt.Errorf("No Mock Found for %s", url)
		}
		return mock.Response, mock.Err
	}
	jsonBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = header

	client := http.Client{}

	return client.Do(request)
}
