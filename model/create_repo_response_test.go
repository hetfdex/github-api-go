package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateRepoResponseDto(t *testing.T) {
	req := CreateRepoResponse{
		ID:   0,
		Name: "name",
	}

	result := req.CreateRepoResponseDto()

	assert.NotNil(t, result)
	assert.EqualValues(t, req.ID, result.ID)
	assert.EqualValues(t, req.Name, result.Name)
}

func TestNewCreateRepoResponseFromBytesFailed(t *testing.T) {
	var body []byte

	result, err := NewCreateRepoResponseFromBytes(body)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "unexpected end of JSON input", err.Message)
}

func TestNewCreateRepoResponseFromBytesOk(t *testing.T) {
	testResponse := &CreateRepoResponse{
		ID:   0,
		Name: "name",
	}
	body, _ := json.Marshal(testResponse)

	result, err := NewCreateRepoResponseFromBytes(body)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, testResponse.ID, result.ID)
	assert.EqualValues(t, testResponse.Name, result.Name)
}
