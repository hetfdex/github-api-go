package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hetfdex/github-api-go/mock"
	"github.com/hetfdex/github-api-go/model"
	"github.com/hetfdex/github-api-go/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var rec *httptest.ResponseRecorder

var ctx *gin.Context

func TestMain(m *testing.M) {
	Controller = &controller{
		&mock.ServiceRepoCreatorMock{},
	}
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonBodyError(t *testing.T) {
	rec = httptest.NewRecorder()

	ctx, _ = gin.CreateTestContext(rec)

	Controller.CreateRepo(ctx)

	var errDto model.ErrorResponseDto

	_ = json.Unmarshal(rec.Body.Bytes(), &errDto)

	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, http.StatusBadRequest, errDto.StatusCode)
	assert.EqualValues(t, util.InvalidJsonBodyError, errDto.Message)
}

func TestCreateRepoError(t *testing.T) {
	req := model.CreateRepoRequestDto{
		Name:        mock.ControllerCreateRepoError,
		Description: "description",
	}

	reqBytes, _ := json.Marshal(req)

	rec = httptest.NewRecorder()

	ctx, _ = gin.CreateTestContext(rec)

	ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBytes))

	Controller.CreateRepo(ctx)

	var errDto model.ErrorResponseDto

	_ = json.Unmarshal(rec.Body.Bytes(), &errDto)

	assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
	assert.EqualValues(t, http.StatusInternalServerError, errDto.StatusCode)
	assert.EqualValues(t, mock.ControllerCreateRepoError, errDto.Message)
}

func TestCreateRepo(t *testing.T) {
	req := model.CreateRepoRequestDto{
		Name:        "name",
		Description: "description",
	}

	reqBytes, _ := json.Marshal(req)

	rec = httptest.NewRecorder()

	ctx, _ = gin.CreateTestContext(rec)

	ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBytes))

	Controller.CreateRepo(ctx)

	var resDto model.CreateRepoResponseDto

	_ = json.Unmarshal(rec.Body.Bytes(), &resDto)

	assert.EqualValues(t, http.StatusCreated, rec.Code)
	assert.EqualValues(t, 0, resDto.ID)
	assert.EqualValues(t, req.Name, resDto.Name)
}

func TestCreateReposInvalidJsonBodyError(t *testing.T) {
	rec = httptest.NewRecorder()

	ctx, _ = gin.CreateTestContext(rec)

	Controller.CreateRepos(ctx)

	var errDto model.ErrorResponseDto

	_ = json.Unmarshal(rec.Body.Bytes(), &errDto)

	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, http.StatusBadRequest, errDto.StatusCode)
	assert.EqualValues(t, util.InvalidJsonBodyError, errDto.Message)
}

func TestCreateRepos(t *testing.T) {
	reqOne := model.CreateRepoRequestDto{
		Name:        "one",
		Description: "description",
	}

	reqTwo := model.CreateRepoRequestDto{
		Name:        "two",
		Description: "description",
	}

	requestsDto := model.CreateReposRequestDto{
		Requests: []model.CreateRepoRequestDto{
			reqOne,
			reqTwo,
		},
	}

	reqBytes, _ := json.Marshal(requestsDto)

	rec = httptest.NewRecorder()

	ctx, _ = gin.CreateTestContext(rec)

	ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBytes))

	Controller.CreateRepos(ctx)

	var responseDto model.CreateReposResponseDto

	_ = json.Unmarshal(rec.Body.Bytes(), &responseDto)

	assert.EqualValues(t, http.StatusCreated, rec.Code)
	assert.EqualValues(t, http.StatusCreated, responseDto.StatusCode)
	assert.EqualValues(t, 0, responseDto.Responses[0].ID)
	assert.EqualValues(t, reqOne.Name, responseDto.Responses[0].Name)
	assert.EqualValues(t, 1, responseDto.Responses[1].ID)
	assert.EqualValues(t, reqTwo.Name, responseDto.Responses[1].Name)
}
