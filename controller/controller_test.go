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
	assert.EqualValues(t, "owner", resDto.Owner)
}
