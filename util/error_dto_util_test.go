package util

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewNotFoundDtoError(t *testing.T) {
	statusCode := http.StatusNotFound
	message := "NewNotFoundError"

	err := NewNotFoundDtoError(message)

	assert.NotNil(t, err)
	assert.EqualValues(t, statusCode, err.StatusCode)
	assert.EqualValues(t, message, err.Message)
}

func TestNewBadRequestDtoError(t *testing.T) {
	statusCode := http.StatusBadRequest
	message := "NewBadRequestError"

	err := NewBadRequestDtoError(message)

	assert.NotNil(t, err)
	assert.EqualValues(t, statusCode, err.StatusCode)
	assert.EqualValues(t, message, err.Message)
}

func TestNewInternalServerDtoError(t *testing.T) {
	statusCode := http.StatusInternalServerError
	message := "NewInternalServerError"

	err := NewInternalServerDtoError(message)

	assert.NotNil(t, err)
	assert.EqualValues(t, statusCode, err.StatusCode)
	assert.EqualValues(t, message, err.Message)
}

func TestNewDtoErrorFromBytesFailed(t *testing.T) {
	statusCode := http.StatusFailedDependency
	var body []byte

	err := NewDtoErrorFromBytes(statusCode, body)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "unexpected end of JSON input", err.Message)
}

func TestDtoNewErrorFromBytesOk(t *testing.T) {
	statusCode := http.StatusFailedDependency
	message := "test_message"

	testError := newDtoError(statusCode, message)

	body, _ := json.Marshal(testError)

	err := NewDtoErrorFromBytes(statusCode, body)

	assert.NotNil(t, err)
	assert.EqualValues(t, testError.StatusCode, err.StatusCode)
	assert.EqualValues(t, testError.Message, err.Message)
}
