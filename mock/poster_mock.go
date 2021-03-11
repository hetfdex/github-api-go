package mock

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/hetfdex/github-api-go/util"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const PostErrorAuthorizationHeaderValue = "postError"
const HandleResponseReadAllErrorAuthorizationHeaderValue = "handleResponseReadAllError"
const HandleResponseNotOkNewErrorFromBytesErrorAuthorizationHeaderValue = "handleResponseNotOkNewErrorFromBytesError"
const HandleResponseNotOkAuthorizationHeaderValue = "handleResponseNotOk"
const HandleResponseOkNewCreateRepoResponseFromBytesErrorAuthorizationHeaderValue = "handleResponseOkNewCreateRepoResponseFromBytesError"

type PosterMock struct {
	PostFunc func(string, http.Header, *bytes.Reader) (*http.Response, error)
}

func (*PosterMock) Post(_ string, header http.Header, _ *bytes.Reader) (*http.Response, error) {
	postErrorAuthorizationHeaderValue := fmt.Sprintf(util.AuthorizationHeaderValue, PostErrorAuthorizationHeaderValue)
	readAllErrorAuthorizationHeaderValue := fmt.Sprintf(util.AuthorizationHeaderValue, HandleResponseReadAllErrorAuthorizationHeaderValue)
	handleResponseNotOkNewErrorFromBytesErrorAuthorizationHeaderValue := fmt.Sprintf(util.AuthorizationHeaderValue, HandleResponseNotOkNewErrorFromBytesErrorAuthorizationHeaderValue)
	handleResponseNotOkAuthorizationHeaderValue := fmt.Sprintf(util.AuthorizationHeaderValue, HandleResponseNotOkAuthorizationHeaderValue)
	handleResponseOkNewCreateRepoResponseFromBytesErrorAuthorizationHeaderValue := fmt.Sprintf(util.AuthorizationHeaderValue, HandleResponseOkNewCreateRepoResponseFromBytesErrorAuthorizationHeaderValue)

	switch header.Get(util.AuthorizationHeaderKey) {
	case postErrorAuthorizationHeaderValue:
		return postErrorResponse()
	case readAllErrorAuthorizationHeaderValue:
		return readAllErrorResponse()
	case handleResponseNotOkNewErrorFromBytesErrorAuthorizationHeaderValue:
		return handleResponseNotOkNewErrorFromBytesError()
	case handleResponseNotOkAuthorizationHeaderValue:
		return handleResponseNotOk()
	case handleResponseOkNewCreateRepoResponseFromBytesErrorAuthorizationHeaderValue:
		return handleResponseOkNewCreateRepoResponseFromBytesError()
	}
	return handleResponseOk()
}

func postErrorResponse() (*http.Response, error) {
	return nil, errors.New(PostErrorAuthorizationHeaderValue)
}

func readAllErrorResponse() (*http.Response, error) {
	body, _ := os.Open("-asf3")

	return &http.Response{
		StatusCode: http.StatusFailedDependency,
		Body:       body,
	}, nil
}

func handleResponseNotOkNewErrorFromBytesError() (*http.Response, error) {
	body := ioutil.NopCloser(strings.NewReader(`{
		"message": 0
		}`))

	return &http.Response{
		StatusCode: http.StatusFailedDependency,
		Body:       body,
	}, nil
}

func handleResponseNotOk() (*http.Response, error) {
	body := ioutil.NopCloser(strings.NewReader(`{
		"message": "message",
		"documentation_url": "documentation_url"
		}`))

	return &http.Response{
		StatusCode: http.StatusFailedDependency,
		Body:       body,
	}, nil
}

func handleResponseOkNewCreateRepoResponseFromBytesError() (*http.Response, error) {
	body := ioutil.NopCloser(strings.NewReader(`{
		"id": 0"
		}`))

	return &http.Response{
		StatusCode: http.StatusCreated,
		Body:       body,
	}, nil
}

func handleResponseOk() (*http.Response, error) {
	body := ioutil.NopCloser(strings.NewReader(`{
		"id": 0,
		"name": "name",
		"full_name": "full_name",
		"owner": {
			"login": "login",
			"id": 0,
			"url": "url",
			"html_url": "html_url"
			},
		"permissions": {
			"admin": true,
			"push": true,
			"pull": true
			}
		}`))

	return &http.Response{
		StatusCode: http.StatusCreated,
		Body:       body,
	}, nil
}
