package provider

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hetfdex/github-api-go/src/api/client"
	"github.com/hetfdex/github-api-go/src/api/model"
)

func TestMain(m *testing.M) {
	client.StartMocker()

	os.Exit(m.Run())
}
func TestCreateRepoInvalidClientResponse(t *testing.T) {
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
	assert.EqualValues(t, "Invalid Error Response Body", err.Message)
}
