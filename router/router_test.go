package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hetfdex/github-api-go/mock"
	"github.com/hetfdex/github-api-go/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var engine *gin.Engine

var rec *httptest.ResponseRecorder

var req *http.Request

func TestMain(m *testing.M) {
	Router = &router{
		&mock.ControllerRepoCreatorMock{},
		&mock.HealthCheckerMock{},
	}
	engine = Router.Setup()

	os.Exit(m.Run())
}

func TestCheckHealth(t *testing.T) {
	rec = httptest.NewRecorder()

	req = httptest.NewRequest(http.MethodGet, util.HealthCheckPath, nil)

	engine.ServeHTTP(rec, req)

	body := strconv.Itoa(http.StatusAccepted)

	assert.EqualValues(t, http.StatusOK, rec.Code)
	assert.EqualValues(t, body, rec.Body.String())
}

func TestCreateRepo(t *testing.T) {
	rec = httptest.NewRecorder()

	req = httptest.NewRequest(http.MethodPost, util.RepositoriesPath, nil)

	engine.ServeHTTP(rec, req)

	body := strconv.Itoa(http.StatusAccepted)

	assert.EqualValues(t, http.StatusCreated, rec.Code)
	assert.EqualValues(t, body, rec.Body.String())
}
