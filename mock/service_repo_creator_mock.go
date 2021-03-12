package mock

import (
	"github.com/hetfdex/github-api-go/model"
)

type ServiceRepoCreatorMock struct {
	CreateRepoFunc func(model.CreateRepoRequestDto) (*model.CreateRepoResponseDto, *model.ErrorResponseDto)
}

func (*ServiceRepoCreatorMock) CreateRepo(model.CreateRepoRequestDto) (*model.CreateRepoResponseDto, *model.ErrorResponseDto) {
	return nil, nil
}
