package service

import (
	"github.com/hetfdex/github-api-go/mock"
	"github.com/hetfdex/github-api-go/model"
	"github.com/hetfdex/github-api-go/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

var reqDto model.CreateRepoRequestDto

func TestMain(m *testing.M) {
	Service = &service{
		&mock.ProviderRepoCreatorMock{},
	}
	os.Exit(m.Run())
}

func TestCreateRepoInvalidRepoName(t *testing.T) {
	reqDto = model.CreateRepoRequestDto{
		Name:        "",
		Description: "description",
	}

	res, err := Service.CreateRepo(reqDto)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode)
	assert.EqualValues(t, util.InvalidRepoNameError, err.Message)
}

func TestCreateRepoError(t *testing.T) {
	reqDto = model.CreateRepoRequestDto{
		Name:        mock.ServiceCreateRepoError,
		Description: "description",
	}

	res, err := Service.CreateRepo(reqDto)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, mock.ServiceCreateRepoError, err.Message)
}

func TestCreateRepo(t *testing.T) {
	reqDto = model.CreateRepoRequestDto{
		Name:        "name",
		Description: "description",
	}

	res, err := Service.CreateRepo(reqDto)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.EqualValues(t, 0, res.ID)
	assert.EqualValues(t, reqDto.Name, res.Name)
	assert.EqualValues(t, "owner_login", res.Owner)
}
