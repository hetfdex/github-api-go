package mapper

import (
	"github.com/hetfdex/github-api-go/dto"
	"github.com/hetfdex/github-api-go/model"
)

func ToCreateRepoRequestModel(req dto.CreateRepoRequest) model.CreateRepoRequest {
	return model.CreateRepoRequest{
		Name:        req.Name,
		Description: req.Description,
		Private:     false,
	}
}
