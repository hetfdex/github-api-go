package model

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotFoundError(t *testing.T) {
	error := NewNotFoundError("message")

	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusNotFound, error.Status())
	assert.EqualValues(t, "message", error.Message())
	assert.EqualValues(t, "", error.Error())

}

func TestNewInternalServerError(t *testing.T) {
	error := NewInternalServerError("message")

	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusInternalServerError, error.Status())
	assert.EqualValues(t, "message", error.Message())
	assert.EqualValues(t, "", error.Error())
}

func TestNewBadRequestError(t *testing.T) {
	error := NewBadRequestError("message")

	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusBadRequest, error.Status())
	assert.EqualValues(t, "message", error.Message())
	assert.EqualValues(t, "", error.Error())
}
