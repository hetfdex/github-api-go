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

type RepoCreator interface {
	CreateRepo(req github.CreateRepoRequest, token string) (*github.CreateRepoResponse, *github.ErrorResponse)
}

type provider struct {
}

var Provider RepoCreator = &provider{}

func (m *provider) CreateRepo(req github.CreateRepoRequest, token string) (*github.CreateRepoResponse, *github.ErrorResponse) {
	header := makeHeader(token)

	body := makeBody(req)

	res, err := rest.Post(util.CreateRepoUrl, header, body)

	if err != nil {
		return nil, util.NewInternalServerError(err.Error())
	}

	return handleResponse(res.StatusCode, res.Body)
}

func makeHeader(token string) http.Header {
	header := http.Header{}

	key := util.AuthorizationHeaderKey

	value := fmt.Sprintf(util.AuthorizationHeaderValue, token)

	header.Set(key, value)

	return header
}

func makeBody(req github.CreateRepoRequest) *bytes.Reader {
	//Error ignored because it's extremely unlikely to occur
	jsonBytes, _ := json.Marshal(req)

	return bytes.NewReader(jsonBytes)
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

func handleResponseNotOk(statusCode int, bytes []byte) (*github.CreateRepoResponse, *github.ErrorResponse) {
	errorResponse, err := util.NewErrorFromBytes(statusCode, bytes)

	if err != nil {
		return nil, err
	}
	return nil, errorResponse
}

func handleResponseOk(bytes []byte) (*github.CreateRepoResponse, *github.ErrorResponse) {
	createRepoResponse, err := util.NewCreateRepoResponseFromBytes(bytes)

	if err != nil {
		return nil, err
	}
	return createRepoResponse, nil
}
