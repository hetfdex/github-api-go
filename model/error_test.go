package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestErrorResponseDto(t *testing.T) {
	err := newError(http.StatusFailedDependency, "message")

	result := err.ErrorResponseDto()

	assert.NotNil(t, result)
	assert.EqualValues(t, err.StatusCode, result.StatusCode)
	assert.EqualValues(t, err.Message, result.Message)
}

func TestNewInternalServerError(t *testing.T) {
	statusCode := http.StatusInternalServerError
	message := "NewNotFoundError"

	err := NewInternalServerError(message)

	assert.NotNil(t, err)
	assert.EqualValues(t, statusCode, err.StatusCode)
	assert.EqualValues(t, message, err.Message)
}

func TestNewErrorFromBytesFailed(t *testing.T) {
	statusCode := http.StatusFailedDependency
	var body []byte

	err := NewErrorFromBytes(statusCode, body)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "unexpected end of JSON input", err.Message)
}

func TestNewErrorFromBytesOk(t *testing.T) {
	statusCode := http.StatusFailedDependency
	message := "test_message"

	testError := newDtoError(statusCode, message)

	body, _ := json.Marshal(testError)

	err := NewErrorFromBytes(statusCode, body)

	assert.NotNil(t, err)
	assert.EqualValues(t, testError.StatusCode, err.StatusCode)
	assert.EqualValues(t, testError.Message, err.Message)
}
