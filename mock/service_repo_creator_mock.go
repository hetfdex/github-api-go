package mock

import (
	"github.com/hetfdex/github-api-go/model"
	"net/http"
)

const ControllerCreateRepoError = "controllerCreateRepoError"

type ServiceRepoCreatorMock struct {
	CreateRepoFunc  func(model.CreateRepoRequestDto) (*model.CreateRepoResponseDto, *model.ErrorResponseDto)
	CreateReposFunc func(model.CreateReposRequestDto) *model.CreateReposResponseDto
}

func (*ServiceRepoCreatorMock) CreateRepo(reqDto model.CreateRepoRequestDto) (*model.CreateRepoResponseDto, *model.ErrorResponseDto) {
	if reqDto.Name == ControllerCreateRepoError {
		return nil, model.NewInternalServerErrorDto(ControllerCreateRepoError)
	}
	return &model.CreateRepoResponseDto{
		ID:   0,
		Name: reqDto.Name,
	}, nil
}

func (*ServiceRepoCreatorMock) CreateRepos(reqsDto model.CreateReposRequestDto) *model.CreateReposResponseDto {
	return &model.CreateReposResponseDto{
		StatusCode: http.StatusCreated,
		Responses: []*model.CreateRepoResponseDto{
			{
				ID:   0,
				Name: reqsDto.Requests[0].Name,
			},
		},
	}
}
