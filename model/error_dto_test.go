package model

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerDtoError(t *testing.T) {
	statusCode := http.StatusInternalServerError
	message := "NewInternalServerError"

	err := NewInternalServerDtoError(message)

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
