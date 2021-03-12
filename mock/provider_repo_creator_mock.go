package mock

import (
	"github.com/hetfdex/github-api-go/model"
)

type ProviderRepoCreatorMock struct {
	CreateRepoFunc func(model.CreateRepoRequest, string) (*model.CreateRepoResponse, *model.ErrorResponse)
}

func (*ProviderRepoCreatorMock) CreateRepo(model.CreateRepoRequest, string) (*model.CreateRepoResponse, *model.ErrorResponse) {
	return nil, nil
}
