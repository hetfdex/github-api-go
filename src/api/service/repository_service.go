package service

import (
	"strings"

	"github.com/hetfdex/github-api-go/src/api/config"
	"github.com/hetfdex/github-api-go/src/api/model"
	"github.com/hetfdex/github-api-go/src/api/provider"
)

type repositoryService struct{}

type repositoryServiceInterface interface {
	CreateRepo(request model.CreateRepoRequest) (*model.CreateRepoResponse, error)
}

// RepositoryService entry point
var RepositoryService repositoryServiceInterface

func init() {
	var RepositoryService = &repositoryService{}
}

func (r *repositoryService) CreateRepo(request model.CreateRepoRequest) (*model.CreateRepoResponse, model.APIError) {
	request.Name = strings.TrimSpace(request.Name)

	if request.Name == "" {
		return nil, model.NewBadRequestError("Invalid Repository Name")
	}
	req := model.GitHubCreateRepoRequest{
		Name:        request.Name,
		Description: request.Description,
		Private:     false,
	}
	res, err := provider.CreateRepo(config.GetGitHubAccessToken(), req)

	if err != nil {
		return nil, model.NewAPIError(err.StatusCode, err.Message)
	}
	response := &model.CreateRepoResponse{
		ID:    res.ID,
		Name:  res.Name,
		Owner: res.Owner.Login,
	}
	return response, nil
}
