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
	authorizationHeaderValue := header.Get(util.AuthorizationHeaderKey)

	postErrorAuthorizationHeaderValue := fmt.Sprintf(util.AuthorizationHeaderValue, PostErrorAuthorizationHeaderValue)
	readAllErrorAuthorizationHeaderValue := fmt.Sprintf(util.AuthorizationHeaderValue, HandleResponseReadAllErrorAuthorizationHeaderValue)
	handleResponseNotOkNewErrorFromBytesErrorAuthorizationHeaderValue := fmt.Sprintf(util.AuthorizationHeaderValue, HandleResponseNotOkNewErrorFromBytesErrorAuthorizationHeaderValue)
	handleResponseNotOkAuthorizationHeaderValue := fmt.Sprintf(util.AuthorizationHeaderValue, HandleResponseNotOkAuthorizationHeaderValue)
	handleResponseOkNewCreateRepoResponseFromBytesErrorAuthorizationHeaderValue := fmt.Sprintf(util.AuthorizationHeaderValue, HandleResponseOkNewCreateRepoResponseFromBytesErrorAuthorizationHeaderValue)

	if authorizationHeaderValue == postErrorAuthorizationHeaderValue {
		return nil, errors.New(PostErrorAuthorizationHeaderValue)
	} else if authorizationHeaderValue == readAllErrorAuthorizationHeaderValue {
		body, _ := os.Open("-asf3")

		return &http.Response{
			StatusCode: http.StatusFailedDependency,
			Body:       body,
		}, nil
	} else if authorizationHeaderValue == handleResponseNotOkNewErrorFromBytesErrorAuthorizationHeaderValue {
		body := ioutil.NopCloser(strings.NewReader(`{
		"message": 0
		}`))

		return &http.Response{
			StatusCode: http.StatusFailedDependency,
			Body:       body,
		}, nil
	} else if authorizationHeaderValue == handleResponseNotOkAuthorizationHeaderValue {
		body := ioutil.NopCloser(strings.NewReader(`{
		"message": "message",
		"documentation_url": "documentation_url"
		}`))

		return &http.Response{
			StatusCode: http.StatusFailedDependency,
			Body:       body,
		}, nil
	} else if authorizationHeaderValue == handleResponseOkNewCreateRepoResponseFromBytesErrorAuthorizationHeaderValue {
		body := ioutil.NopCloser(strings.NewReader(`{
		"id": 0"
		}`))

		return &http.Response{
			StatusCode: http.StatusCreated,
			Body:       body,
		}, nil
	}
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
