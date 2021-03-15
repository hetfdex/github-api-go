package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCreateReposResponseChan(t *testing.T) {
	res := CreateRepoResponse{
		ID:   0,
		Name: "name",
		Owner: Owner{
			Login: "login",
		},
	}

	err := ErrorResponse{
		StatusCode: 1,
		Message:    "message",
	}

	result := NewCreateReposResponseChan(res, err)

	assert.NotNil(t, result)
	assert.EqualValues(t, result.Response.ID, res.ID)
	assert.EqualValues(t, result.Response.Name, res.Name)
	assert.EqualValues(t, result.Response.Owner, res.Owner.Login)
	assert.EqualValues(t, result.Error.StatusCode, err.StatusCode)
	assert.EqualValues(t, result.Error.Message, err.Message)
}
