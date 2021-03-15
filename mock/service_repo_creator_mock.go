package mock

import (
	"github.com/hetfdex/github-api-go/model"
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
		Name: "name",
	}, nil
}

func (*ServiceRepoCreatorMock) CreateRepos(model.CreateReposRequestDto) *model.CreateReposResponseDto {
	return nil
}
