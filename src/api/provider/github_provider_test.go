package provider

import (
	"errors"
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
func TestCreateRepoInvalidResponse(t *testing.T) {
	message := "Invalid Client Response"

	mock := client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Err:        errors.New(message),
	}
	client.FlushMocks()

	client.AddMock(mock)

	response, err := CreateRepo("", model.GitHubCreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, message, err.Message)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	invalidResponseBody, _ := os.Open("-asf3")

	response := &http.Response{
		StatusCode: http.StatusCreated,
		Body:       invalidResponseBody,
	}

	mock := client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response:   response,
	}
	client.FlushMocks()

	client.AddMock(mock)

	res, err := CreateRepo("", model.GitHubCreateRepoRequest{})

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid Response Body", err.Message)
}

func TestCreateRepoInvalidErrorResponse(t *testing.T) {
	invalidErrorResponse := ioutil.NopCloser(strings.NewReader(`{
		"message": 0
		}`))

	response := &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       invalidErrorResponse,
	}

	mock := client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response:   response,
	}
	client.FlushMocks()

	client.AddMock(mock)

	res, err := CreateRepo("", model.GitHubCreateRepoRequest{})

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid Error Response", err.Message)
}

func TestCreateRepoValidErrorResponse(t *testing.T) {
	validErrorResponse := ioutil.NopCloser(strings.NewReader(`{
		"message": "message",
		"documentation_url": "documentation_url"
		}`))

	response := &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       validErrorResponse,
	}

	mock := client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response:   response,
	}
	client.FlushMocks()

	client.AddMock(mock)

	res, err := CreateRepo("", model.GitHubCreateRepoRequest{})

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "message", err.Message)
	assert.EqualValues(t, "documentation_url", err.DocumentationURL)
}
func TestCreateRepoInvalidCreatedResponse(t *testing.T) {
	invalidCreatedResponse := ioutil.NopCloser(strings.NewReader(`{
		"id": 0"
		}`))

	response := &http.Response{
		StatusCode: http.StatusCreated,
		Body:       invalidCreatedResponse,
	}

	mock := client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response:   response,
	}
	client.FlushMocks()

	client.AddMock(mock)

	res, err := CreateRepo("", model.GitHubCreateRepoRequest{})

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid Created Response", err.Message)
}
