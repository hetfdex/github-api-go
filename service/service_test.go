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

func TestMain(m *testing.M) {
	Service = &service{
		&mock.ProviderRepoCreatorMock{},
	}
	os.Exit(m.Run())
}

func TestCreateRepoInvalidRepoName(t *testing.T) {
	reqDto := model.CreateRepoRequestDto{
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
	reqDto := model.CreateRepoRequestDto{
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
	reqDto := model.CreateRepoRequestDto{
		Name:        "name",
		Description: "description",
	}

	res, err := Service.CreateRepo(reqDto)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.EqualValues(t, 666, res.ID)
	assert.EqualValues(t, reqDto.Name, res.Name)
}

func TestCreateReposError(t *testing.T) {
	reqDtoOne := model.CreateRepoRequestDto{
		Name:        mock.ServiceCreateRepoError,
		Description: "description",
	}

	reqDtoTwo := model.CreateRepoRequestDto{
		Name:        mock.ServiceCreateRepoError,
		Description: "description",
	}

	requestsDto := model.CreateReposRequestDto{
		Requests: []model.CreateRepoRequestDto{
			reqDtoOne,
			reqDtoTwo,
		},
	}

	res := Service.CreateRepos(requestsDto)

	assert.NotNil(t, res)
	assert.Nil(t, res.Responses)
	assert.NotNil(t, res.Errors)
	assert.EqualValues(t, http.StatusInternalServerError, res.StatusCode)
	assert.EqualValues(t, len(requestsDto.Requests), len(res.Errors))
}

func TestCreateReposMixed(t *testing.T) {
	reqDtoOne := model.CreateRepoRequestDto{
		Name:        "name",
		Description: "description",
	}

	reqDtoTwo := model.CreateRepoRequestDto{
		Name:        mock.ServiceCreateRepoError,
		Description: "description",
	}

	requestsDto := model.CreateReposRequestDto{
		Requests: []model.CreateRepoRequestDto{
			reqDtoOne,
			reqDtoTwo,
		},
	}

	res := Service.CreateRepos(requestsDto)

	assert.NotNil(t, res)
	assert.NotNil(t, res.Responses)
	assert.NotNil(t, res.Errors)
	assert.EqualValues(t, http.StatusPartialContent, res.StatusCode)
	assert.EqualValues(t, 1, len(res.Responses))
	assert.EqualValues(t, 1, len(res.Errors))
}

func TestCreateRepos(t *testing.T) {
	reqDtoOne := model.CreateRepoRequestDto{
		Name:        "name",
		Description: "description",
	}

	reqDtoTwo := model.CreateRepoRequestDto{
		Name:        "name",
		Description: "description",
	}

	requestsDto := model.CreateReposRequestDto{
		Requests: []model.CreateRepoRequestDto{
			reqDtoOne,
			reqDtoTwo,
		},
	}

	res := Service.CreateRepos(requestsDto)

	assert.NotNil(t, res)
	assert.Nil(t, res.Errors)
	assert.NotNil(t, res.Responses)
	assert.EqualValues(t, http.StatusCreated, res.StatusCode)
	assert.EqualValues(t, len(requestsDto.Requests), len(res.Responses))
}
