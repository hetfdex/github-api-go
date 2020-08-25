package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/hetfdex/github-api-go/src/api/client"

	"github.com/hetfdex/github-api-go/src/api/model"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	client.StartMocker()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidBody(t *testing.T) {
	response := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(response)

	request, _ := http.NewRequest(http.MethodPost, "/url", strings.NewReader(``))

	c.Request = request

	CreateRepo(c)

	apiError, err := model.NewAPIErrorFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.EqualValues(t, http.StatusBadRequest, apiError.GetStatusCode())
	assert.EqualValues(t, "Invalid JSON Body", apiError.GetMessage())
}

func TestCreateRepoServiceError(t *testing.T) {
	responseBody := ioutil.NopCloser(strings.NewReader(`{
		"message":"message"
		}`))

	response := &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       responseBody,
	}

	mock := client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response:   response,
	}
	client.FlushMocks()

	client.AddMock(mock)

	res := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(res)

	body := `{
		"name": "name"
		}`

	request, _ := http.NewRequest(http.MethodPost, "/url", strings.NewReader(body))

	c.Request = request

	CreateRepo(c)

	apiError, err := model.NewAPIErrorFromBytes(res.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, http.StatusUnauthorized, res.Code)
	assert.EqualValues(t, http.StatusUnauthorized, apiError.GetStatusCode())
	assert.EqualValues(t, "message", apiError.GetMessage())
}

func TestCreateRepoValidResponse(t *testing.T) {
	validCreatedResponse := ioutil.NopCloser(strings.NewReader(`{
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

	response := &http.Response{
		StatusCode: http.StatusCreated,
		Body:       validCreatedResponse,
	}

	mock := client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response:   response,
	}
	client.FlushMocks()

	client.AddMock(mock)

	res := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(res)

	body := `{
		"name": "name",
		"description":"description"
		}`

	request, _ := http.NewRequest(http.MethodPost, "/url", strings.NewReader(body))

	c.Request = request

	CreateRepo(c)

	var validResponse model.CreateRepoResponse

	err := json.Unmarshal(res.Body.Bytes(), &validResponse)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusCreated, res.Code)
	assert.EqualValues(t, 0, validResponse.ID)
	assert.EqualValues(t, "name", validResponse.Name)
	assert.EqualValues(t, "login", validResponse.Owner)
}