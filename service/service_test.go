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
var reqDtoConcurrent model.CreateRepoRequestDto
var reqsDto model.CreateReposRequestDto

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
}

func TestCreateReposError(t *testing.T) {
	reqDto = model.CreateRepoRequestDto{
		Name:        mock.ServiceCreateRepoError,
		Description: "description",
	}

	reqDtoConcurrent = model.CreateRepoRequestDto{
		Name:        mock.ServiceCreateRepoError,
		Description: "description",
	}

	reqsDto = model.CreateReposRequestDto{
		Requests: []model.CreateRepoRequestDto{
			reqDto,
			reqDtoConcurrent,
		},
	}

	res := Service.CreateRepos(reqsDto)

	assert.NotNil(t, res)
	assert.NotNil(t, res.Responses)
	assert.NotNil(t, res.Errors)
	assert.EqualValues(t, http.StatusInternalServerError, res.StatusCode)
}

func TestCreateReposMixed(t *testing.T) {
	reqDto = model.CreateRepoRequestDto{
		Name:        "name",
		Description: "description",
	}

	reqDtoConcurrent = model.CreateRepoRequestDto{
		Name:        mock.ServiceCreateRepoError,
		Description: "description",
	}

	reqsDto = model.CreateReposRequestDto{
		Requests: []model.CreateRepoRequestDto{
			reqDto,
			reqDtoConcurrent,
		},
	}

	res := Service.CreateRepos(reqsDto)

	assert.NotNil(t, res)
	assert.NotNil(t, res.Responses)
	assert.NotNil(t, res.Errors)
	assert.EqualValues(t, http.StatusPartialContent, res.StatusCode)
}

func TestCreateRepos(t *testing.T) {
	reqDto = model.CreateRepoRequestDto{
		Name:        "name",
		Description: "description",
	}

	reqDtoConcurrent = model.CreateRepoRequestDto{
		Name:        "name",
		Description: "description",
	}

	reqsDto = model.CreateReposRequestDto{
		Requests: []model.CreateRepoRequestDto{
			reqDto,
			reqDtoConcurrent,
		},
	}

	res := Service.CreateRepos(reqsDto)

	assert.NotNil(t, res)
	assert.NotNil(t, res.Responses)
	assert.NotNil(t, res.Errors)
	assert.EqualValues(t, http.StatusCreated, res.StatusCode)
}
