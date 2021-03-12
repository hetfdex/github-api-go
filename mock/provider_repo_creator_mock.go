package mock

import (
	"github.com/hetfdex/github-api-go/model"
)

const ServiceCreateRepoError = "serviceCreateRepoError"

type ProviderRepoCreatorMock struct {
	CreateRepoFunc func(model.CreateRepoRequest, string) (*model.CreateRepoResponse, *model.ErrorResponse)
}

func (*ProviderRepoCreatorMock) CreateRepo(req model.CreateRepoRequest, _ string) (*model.CreateRepoResponse, *model.ErrorResponse) {
	if req.Name == ServiceCreateRepoError {
		return nil, model.NewInternalServerError(ServiceCreateRepoError)
	}
	return &model.CreateRepoResponse{
		ID:       0,
		Name:     "name",
		FullName: "fullName",
		Owner: model.Owner{
			Login:   "owner_login",
			ID:      0,
			URL:     "url",
			HtmlUrl: "htmlUrl",
		},
		Permissions: model.Permissions{
			Admin: false,
			Push:  false,
			Pull:  false,
		},
	}, nil
}
