package model

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const message = "message"

func TestNewAPIError(t *testing.T) {
	const statusCode = 424

	error := NewAPIError(statusCode, message)

	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusFailedDependency, error.GetStatusCode())
	assert.EqualValues(t, message, error.GetMessage())
	assert.EqualValues(t, "", error.GetError())

}

func TestNewNotFoundError(t *testing.T) {
	error := NewNotFoundError(message)

	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusNotFound, error.GetStatusCode())
	assert.EqualValues(t, message, error.GetMessage())
	assert.EqualValues(t, "", error.GetError())

}

func TestNewInternalServerError(t *testing.T) {
	error := NewInternalServerError(message)

	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusInternalServerError, error.GetStatusCode())
	assert.EqualValues(t, message, error.GetMessage())
	assert.EqualValues(t, "", error.GetError())
}

func TestNewBadRequestError(t *testing.T) {
	error := NewBadRequestError(message)

	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusBadRequest, error.GetStatusCode())
	assert.EqualValues(t, message, error.GetMessage())
	assert.EqualValues(t, "", error.GetError())
}
