package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hetfdex/github-api-go/client/rest"
	"github.com/hetfdex/github-api-go/model/github"
	"github.com/hetfdex/github-api-go/util"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateRepo(req github.CreateRepoRequest, token string) (*github.CreateRepoResponse, *github.ErrorResponse) {
	header := makeHeader(token)

	body, err := makeBody(req)

	if err != nil {
		return nil, util.NewInternalServerError(err.Error())
	}

	res, err := rest.Post(util.CreateRepoUrl, header, body)

	if err != nil {
		return nil, util.NewInternalServerError(err.Error())
	}

	return handleResponse(res.StatusCode, res.Body)
}

func makeHeader(token string) http.Header {
	header := http.Header{}

	header.Set(util.AuthorizationHeaderKey, fmt.Sprintf(util.AuthorizationHeaderValue, token))

	return header
}

func makeBody(body interface{}) (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}
	return bytes.NewReader(jsonBytes), nil
}

func handleResponse(statusCode int, body io.ReadCloser) (*github.CreateRepoResponse, *github.ErrorResponse) {
	bodyBytes, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, util.NewInternalServerError(err.Error())
	}

	defer func() {
		_ = body.Close()
	}()

	if statusCode > 299 {
		return handleResponseNotOk(statusCode, bodyBytes)
	}
	return handleResponseOk(bodyBytes)

}
func handleResponseOk(bytes []byte) (*github.CreateRepoResponse, *github.ErrorResponse) {
	createRepoResponse, err := util.NewCreateRepoFromBytes(bytes)

	if err != nil {
		return nil, err
	}
	return createRepoResponse, nil
}

func handleResponseNotOk(statusCode int, bytes []byte) (*github.CreateRepoResponse, *github.ErrorResponse) {
	errorResponse, err := util.NewErrorFromBytes(statusCode, bytes)

	if err != nil {
		return nil, err
	}
	return nil, errorResponse
}
