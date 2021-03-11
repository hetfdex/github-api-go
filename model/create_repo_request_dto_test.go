package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequest(t *testing.T) {
	req := CreateRepoRequestDto{
		Name:        "name",
		Description: "description",
	}

	result := req.CreateRepoRequest()

	assert.NotNil(t, result)
	assert.EqualValues(t, req.Name, result.Name)
	assert.EqualValues(t, req.Description, result.Description)
}
