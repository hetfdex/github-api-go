package mock

import (
	"errors"
	"net/http"
)

const DoErrorUrl = "doErrorUrl"

type DoerMock struct {
	DoFunc func(*http.Request) (*http.Response, error)
}

func (*DoerMock) Do(req *http.Request) (*http.Response, error) {
	if req.URL.String() == DoErrorUrl {
		return nil, errors.New(DoErrorUrl)
	}
	return &http.Response{
		StatusCode: http.StatusOK,
	}, nil
}
