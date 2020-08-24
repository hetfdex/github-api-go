package service

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hetfdex/github-api-go/src/api/client"
	"github.com/hetfdex/github-api-go/src/api/model"
)

func TestMain(m *testing.M) {
	client.StartMocker()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidName(t *testing.T) {
	request := model.CreateRepoRequest{}

	response, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.GetStatusCode())
	assert.EqualValues(t, "Invalid Repository Name", err.GetMessage())
}

func TestCreateRepoProviderError(t *testing.T) {
	body := ioutil.NopCloser(strings.NewReader(`{
		"message": "message",
		"documentation_url": "documentation_url"
		}`))

	response := &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       body,
	}

	mock := client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response:   response,
	}
	client.FlushMocks()

	client.AddMock(mock)

	request := model.CreateRepoRequest{
		Name:        "name",
		Description: "description",
	}

	res, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.GetStatusCode())
	assert.EqualValues(t, "message", err.GetMessage())
}

func TestCreateRepoValidResponse(t *testing.T) {
	validCreatedResponse := ioutil.NopCloser(strings.NewReader(`{
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

	response := &http.Response{
		StatusCode: http.StatusCreated,
		Body:       validCreatedResponse,
	}

	mock := client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response:   response,
	}
	client.FlushMocks()

	client.AddMock(mock)

	request := model.CreateRepoRequest{
		Name:        "name",
		Description: "description",
	}

	res, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.EqualValues(t, 0, res.ID)
	assert.EqualValues(t, "name", res.Name)
	assert.EqualValues(t, "login", res.Owner)
}
