package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hetfdex/github-api-go/client"
	"github.com/hetfdex/github-api-go/model"
	"github.com/hetfdex/github-api-go/util"
	"io"
	"io/ioutil"
	"net/http"
)

type RepoCreator interface {
	CreateRepo(model.CreateRepoRequest, string) (*model.CreateRepoResponse, *model.ErrorResponse)
}

type provider struct {
	client.Poster
}

var Provider RepoCreator = &provider{
	client.PostClient,
}

func (p *provider) CreateRepo(req model.CreateRepoRequest, token string) (*model.CreateRepoResponse, *model.ErrorResponse) {
	header := makeHeader(token)
	body := makeBody(req)

	res, err := p.Post(util.CreateRepoUrl, header, body)

	if err != nil {
		return nil, model.NewInternalServerError(err.Error())
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

func makeBody(req model.CreateRepoRequest) *bytes.Reader {
	//Error ignored because it's extremely unlikely to occur
	jsonBytes, _ := json.Marshal(req)

	return bytes.NewReader(jsonBytes)
}

func handleResponse(statusCode int, body io.ReadCloser) (*model.CreateRepoResponse, *model.ErrorResponse) {
	defer body.Close()

	bodyBytes, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, model.NewInternalServerError(err.Error())
	}

	if statusCode > 299 {
		return nil, handleResponseNotOk(statusCode, bodyBytes)
	}
	return handleResponseOk(bodyBytes)

}

func handleResponseNotOk(statusCode int, bytes []byte) *model.ErrorResponse {
	return model.NewErrorFromBytes(statusCode, bytes)
}

func handleResponseOk(bytes []byte) (*model.CreateRepoResponse, *model.ErrorResponse) {
	createRepoResponse, err := model.NewCreateRepoResponseFromBytes(bytes)

	if err != nil {
		return nil, err
	}
	return createRepoResponse, nil
}
