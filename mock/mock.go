package mock

import (
	"errors"
	"net/http"
)

type DoerMock struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

const DoErrorUrl = "doErrorUrl"

func (m *DoerMock) Do(req *http.Request) (*http.Response, error) {
	if req.URL.String() == DoErrorUrl {
		return nil, errors.New(DoErrorUrl)
	}
	return &http.Response{
		StatusCode: 200,
	}, nil
}
