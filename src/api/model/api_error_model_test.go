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
	assert.EqualValues(t, http.StatusFailedDependency, error.StatusCode())
	assert.EqualValues(t, message, error.Message())
	assert.EqualValues(t, "", error.Error())

}

func TestNewNotFoundError(t *testing.T) {
	error := NewNotFoundError(message)

	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusNotFound, error.StatusCode())
	assert.EqualValues(t, message, error.Message())
	assert.EqualValues(t, "", error.Error())

}

func TestNewInternalServerError(t *testing.T) {
	error := NewInternalServerError(message)

	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusInternalServerError, error.StatusCode())
	assert.EqualValues(t, message, error.Message())
	assert.EqualValues(t, "", error.Error())
}

func TestNewBadRequestError(t *testing.T) {
	error := NewBadRequestError(message)

	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusBadRequest, error.StatusCode())
	assert.EqualValues(t, message, error.Message())
	assert.EqualValues(t, "", error.Error())
}
