package github

import (
	"bytes"
	"errors"
	"github.com/hetfdex/github-api-go/client/rest"
	"github.com/hetfdex/github-api-go/model/github"
	"github.com/hetfdex/github-api-go/util"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestCreateRepoPostError(t *testing.T) {
	mock := &rest.Mock{
		Url:    util.CreateRepoUrl,
		Method: http.MethodPost,
		Err:    errors.New("invalid response"),
	}
	rest.StartMock(mock)

	var req github.CreateRepoRequest

	res, err := Provider.CreateRepo(req, "")

	rest.StopMock()

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, mock.Err.Error(), err.Message)
}

func TestCreateRepoHandleResponseReadAllError(t *testing.T) {
	body, _ := os.Open("-asf3")

	mock := &rest.Mock{
		Url:    util.CreateRepoUrl,
		Method: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusFailedDependency,
			Body:       body,
		},
	}
	rest.StartMock(mock)

	var req github.CreateRepoRequest

	res, err := Provider.CreateRepo(req, "")

	rest.StopMock()

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid argument", err.Message)
}

func TestCreateRepoHandleResponseNotOkNewErrorFromBytesError(t *testing.T) {
	mock := &rest.Mock{
		Url:    util.CreateRepoUrl,
		Method: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusFailedDependency,
			Body:       ioutil.NopCloser(bytes.NewBufferString("bla")),
		},
	}
	rest.StartMock(mock)

	var req github.CreateRepoRequest

	res, err := Provider.CreateRepo(req, "")

	rest.StopMock()

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, util.UnmarshalJsonFailureMessage, err.Message)
}

func TestCreateRepoHandleResponseNotOk(t *testing.T) {
	body := ioutil.NopCloser(strings.NewReader(`{
		"message": "message",
		"documentation_url": "documentation_url"
		}`))

	mock := &rest.Mock{
		Url:    util.CreateRepoUrl,
		Method: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusFailedDependency,
			Body:       body,
		},
	}
	rest.StartMock(mock)

	var req github.CreateRepoRequest

	res, err := Provider.CreateRepo(req, "")

	rest.StopMock()

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusFailedDependency, err.StatusCode)
	assert.EqualValues(t, "message", err.Message)
	assert.EqualValues(t, "documentation_url", err.DocumentationURL)
}

func TestCreateRepoHandleResponseOkNewCreateRepoResponseFromBytesError(t *testing.T) {
	mock := &rest.Mock{
		Url:    util.CreateRepoUrl,
		Method: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(bytes.NewBufferString("bla")),
		},
	}
	rest.StartMock(mock)

	var req github.CreateRepoRequest

	res, err := Provider.CreateRepo(req, "")

	rest.StopMock()

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, util.UnmarshalJsonFailureMessage, err.Message)
}

func TestCreateRepoHandleResponseOk(t *testing.T) {
	body := ioutil.NopCloser(strings.NewReader(`{
		"id": 0,
		"name": "name",
		"full_name": "full_name",
		"owner": {
			"login": "login",
			"id": 0,
			"url": "url",
			"html_url": "html_url"
			},
		"permissions": {
			"admin": true,
			"push": true,
			"pull": true
			}
		}`))

	mock := &rest.Mock{
		Url:    util.CreateRepoUrl,
		Method: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       body,
		},
	}
	rest.StartMock(mock)

	var req github.CreateRepoRequest

	res, err := Provider.CreateRepo(req, "")

	rest.StopMock()

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
