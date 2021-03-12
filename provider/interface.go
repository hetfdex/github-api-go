package provider

import (
	"github.com/hetfdex/github-api-go/client"
	"github.com/hetfdex/github-api-go/model"
)

type RepoCreator interface {
	CreateRepo(model.CreateRepoRequest, string) (*model.CreateRepoResponse, *model.ErrorResponse)
}

type provider struct {
	client.Poster
}
