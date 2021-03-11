package provider

import (
	"github.com/hetfdex/github-api-go/mock"
	"github.com/hetfdex/github-api-go/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

var req model.CreateRepoRequest

func TestMain(m *testing.M) {
	Provider = &provider{
		&mock.PosterMock{},
	}
	os.Exit(m.Run())
}

func TestCreateRepoPostError(t *testing.T) {
	res, err := Provider.CreateRepo(req, mock.PostErrorAuthorizationHeaderValue)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, mock.PostErrorAuthorizationHeaderValue, err.Message)
}

func TestCreateRepoHandleResponseReadAllError(t *testing.T) {
	res, err := Provider.CreateRepo(req, mock.HandleResponseReadAllErrorAuthorizationHeaderValue)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid argument", err.Message)
}

func TestCreateRepoHandleResponseNotOkNewErrorFromBytesError(t *testing.T) {
	res, err := Provider.CreateRepo(req, mock.HandleResponseNotOkNewErrorFromBytesErrorAuthorizationHeaderValue)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "json: cannot unmarshal number into Go struct field ErrorResponse.message of type string", err.Message)
}

func TestCreateRepoHandleResponseNotOk(t *testing.T) {
	res, err := Provider.CreateRepo(req, mock.HandleResponseNotOkAuthorizationHeaderValue)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusFailedDependency, err.StatusCode)
	assert.EqualValues(t, "message", err.Message)
	assert.EqualValues(t, "documentation_url", err.DocumentationURL)
}

func TestCreateRepoHandleResponseOkNewCreateRepoResponseFromBytesError(t *testing.T) {
	res, err := Provider.CreateRepo(req, mock.HandleResponseOkNewCreateRepoResponseFromBytesErrorAuthorizationHeaderValue)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid character '\"' after object key:value pair", err.Message)
}

func TestCreateRepoHandleResponseOk(t *testing.T) {
	res, err := Provider.CreateRepo(req, "")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.EqualValues(t, 0, res.ID)
	assert.EqualValues(t, "name", res.Name)
	assert.EqualValues(t, "full_name", res.FullName)
	assert.EqualValues(t, 0, res.Owner.ID)
	assert.EqualValues(t, "login", res.Owner.Login)
	assert.EqualValues(t, "url", res.Owner.URL)
	assert.EqualValues(t, "html_url", res.Owner.HtmlUrl)
	assert.EqualValues(t, true, res.Permissions.Admin)
	assert.EqualValues(t, true, res.Permissions.Push)
	assert.EqualValues(t, true, res.Permissions.Pull)
}
