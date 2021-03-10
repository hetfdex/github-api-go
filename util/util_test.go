package util

import (
	"encoding/json"
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
