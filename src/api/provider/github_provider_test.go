package provider

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hetfdex/github-api-go/src/api/client"
	"github.com/hetfdex/github-api-go/src/api/model"
)

func TestCreateRepoInvalidClientResponse(t *testing.T) {
	message := "Invalid Client Response"

	mock := client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Err:        errors.New(message),
	}

	client.StartMocker()
	client.AddMock(mock)

	response, err := CreateRepo("", model.GitHubCreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, message, err.Message)
}
