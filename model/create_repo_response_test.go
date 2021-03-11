package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

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
		ID:       0,
		Name:     "name",
		FullName: "fullName",
		Owner: Owner{
			Login:   "login",
			ID:      1,
			URL:     "url",
			HtmlUrl: "htmlUrl",
		},
		Permissions: Permissions{
			Admin: false,
			Push:  false,
			Pull:  false,
		},
	}
	body, _ := json.Marshal(testResponse)

	result, err := NewCreateRepoResponseFromBytes(body)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, testResponse.ID, result.ID)
	assert.EqualValues(t, testResponse.Name, result.Name)
	assert.EqualValues(t, testResponse.FullName, result.FullName)
	assert.EqualValues(t, testResponse.Owner.Login, result.Owner.Login)
	assert.EqualValues(t, testResponse.Owner.ID, result.Owner.ID)
	assert.EqualValues(t, testResponse.Owner.URL, result.Owner.URL)
	assert.EqualValues(t, testResponse.Owner.HtmlUrl, result.Owner.HtmlUrl)
	assert.EqualValues(t, testResponse.Permissions.Admin, result.Permissions.Admin)
	assert.EqualValues(t, testResponse.Permissions.Push, result.Permissions.Push)
	assert.EqualValues(t, testResponse.Permissions.Pull, result.Permissions.Pull)
}
