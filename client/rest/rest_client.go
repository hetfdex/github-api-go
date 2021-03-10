package rest

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

var mocking = false

var mocks map[string]*Mock

type Mock struct {
	Url      string
	Method   string
	Response *http.Response
	Err      error
}

func StartMock(mock *Mock) {
	mocks = make(map[string]*Mock)

	id := getMockId(mock.Method, mock.Url)

	mocks[id] = mock

	mocking = true
}

func StopMock() {
	mocking = false

	mocks = make(map[string]*Mock)
}

func Post(url string, header http.Header, body *bytes.Reader) (*http.Response, error) {
	if mocking {
		return postMock(url)
	}
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

func postMock(url string) (*http.Response, error) {
	id := getMockId(http.MethodPost, url)

	mock := mocks[id]

	if mock == nil {
		message := fmt.Sprintf("no mock available for %v", url)

		err := errors.New(message)

		return nil, err
	}
	return mock.Response, mock.Err
}

func getMockId(method string, url string) string {
	return fmt.Sprintf("%s_%s", method, url)
}
