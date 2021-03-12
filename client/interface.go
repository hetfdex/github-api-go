package client

import (
	"bytes"
	"net/http"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Poster interface {
	Post(string, http.Header, *bytes.Reader) (*http.Response, error)
}

type poster struct {
}
