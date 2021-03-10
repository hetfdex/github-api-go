package util

import (
	"encoding/json"
	"github.com/hetfdex/github-api-go/model/github"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewNotFoundError(t *testing.T) {
	statusCode := http.StatusNotFound

	message := "NewNotFoundError"

	err := NewNotFoundError(message)

	assert.NotNil(t, err)

	assert.EqualValues(t, statusCode, err.StatusCode)

	assert.EqualValues(t, message, err.Message)
}

func TestNewBadRequestError(t *testing.T) {
	statusCode := http.StatusBadRequest

	message := "NewBadRequestError"

	err := NewBadRequestError(message)

	assert.NotNil(t, err)

	assert.EqualValues(t, statusCode, err.StatusCode)

	assert.EqualValues(t, message, err.Message)
}

func TestNewInternalServerError(t *testing.T) {
	statusCode := http.StatusInternalServerError

	message := "NewInternalServerError"

	err := NewInternalServerError(message)

	assert.NotNil(t, err)

	assert.EqualValues(t, statusCode, err.StatusCode)

	assert.EqualValues(t, message, err.Message)
}

func TestNewErrorFromBytesFailed(t *testing.T) {
	statusCode := http.StatusFailedDependency

	var body []byte

	result, err := NewErrorFromBytes(statusCode, body)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, UnmarshalJsonFailureMessage, err.Message)
}

func TestNewErrorFromBytesOk(t *testing.T) {
	statusCode := http.StatusFailedDependency

	message := "test_message"

	testError := newError(http.StatusFailedDependency, message)

	body, _ := json.Marshal(testError)

	result, err := NewErrorFromBytes(statusCode, body)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, testError.StatusCode, result.StatusCode)
	assert.EqualValues(t, testError.Message, result.Message)
}

func TestNewCreateRepoResponseFromBytesFailed(t *testing.T) {
	var body []byte

	result, err := NewCreateRepoResponseFromBytes(body)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, UnmarshalJsonFailureMessage, err.Message)
}

func TestNewCreateRepoResponseFromBytesOk(t *testing.T) {
	testResponse := &github.CreateRepoResponse{
		ID:       0,
		Name:     "name",
		FullName: "fullName",
		Owner: github.Owner{
			Login:   "login",
			ID:      1,
			URL:     "url",
			HtmlUrl: "htmlUrl",
		},
		Permissions: github.Permissions{
			Admin: false,
			Push:  false,
			Pull:  false,
		},
	}

	body, _ := json.Marshal(testResponse)

	result, err := NewCreateRepoResponseFromBytes(body)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, testResponse.ID, result.ID)
	assert.EqualValues(t, testResponse.Name, result.Name)
	assert.EqualValues(t, testResponse.FullName, result.FullName)
	assert.EqualValues(t, testResponse.Owner.Login, result.Owner.Login)
	assert.EqualValues(t, testResponse.Owner.ID, result.Owner.ID)
	assert.EqualValues(t, testResponse.Owner.URL, result.Owner.URL)
	assert.EqualValues(t, testResponse.Owner.HtmlUrl, result.Owner.HtmlUrl)
	assert.EqualValues(t, testResponse.Permissions.Admin, result.Permissions.Admin)
	assert.EqualValues(t, testResponse.Permissions.Push, result.Permissions.Push)
	assert.EqualValues(t, testResponse.Permissions.Pull, result.Permissions.Pull)
}
