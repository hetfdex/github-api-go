package service

import (
	"github.com/hetfdex/github-api-go/config"
	"github.com/hetfdex/github-api-go/model"
	"github.com/hetfdex/github-api-go/provider"
	"github.com/hetfdex/github-api-go/util"
	"strings"
)

var Service RepoCreator = &service{
	provider.Provider,
}

func (s *service) CreateRepo(req model.CreateRepoRequestDto) (*model.CreateRepoResponseDto, *model.ErrorResponseDto) {
	req.Name = strings.TrimSpace(req.Name)

	if req.Name == "" {
		return nil, model.NewBadRequestDtoError(util.InvalidRepoNameError)
	}
	result := req.CreateRepoRequest()

	res, err := s.RepoCreator.CreateRepo(result, config.GetGitHubTokenValue())

	if err != nil {
		dtoErr := err.ErrorResponseDto()

		return nil, dtoErr
	}
	resDto := res.CreateRepoResponseDto()

	return resDto, nil
}
