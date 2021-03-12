package service

import (
	"github.com/hetfdex/github-api-go/model"
	"github.com/hetfdex/github-api-go/provider"
)

type RepoCreator interface {
	CreateRepo(model.CreateRepoRequestDto) (*model.CreateRepoResponseDto, *model.ErrorResponseDto)
	CreateRepos(model.CreateReposRequestDto) (*model.CreateReposResponseDto, *model.ErrorResponseDto)
}

type service struct {
	provider.RepoCreator
}
