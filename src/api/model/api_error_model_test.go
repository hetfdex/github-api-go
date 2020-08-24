package model

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const message = "message"
const statusCode = 424

func TestNewAPIError(t *testing.T) {
	error := NewAPIError(statusCode, message)

	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusFailedDependency, error.GetStatusCode())
	assert.EqualValues(t, message, error.GetMessage())
	assert.EqualValues(t, "", error.GetError())
}

func TestNewAPIErrorFromBytesInvalidBody(t *testing.T) {
	var body []byte

	result, err := NewAPIErrorFromBytes(body)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Invalid JSON Body", err.Error())
}

func TestNewAPIErrorFromBytes(t *testing.T) {
	error := NewAPIError(statusCode, message)

	body, _ := json.Marshal(error)

	result, err := NewAPIErrorFromBytes(body)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, statusCode, result.GetStatusCode())
	assert.EqualValues(t, message, result.GetMessage())
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
