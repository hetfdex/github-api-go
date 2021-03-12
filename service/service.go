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

func (s *service) CreateRepo(reqDto model.CreateRepoRequestDto) (*model.CreateRepoResponseDto, *model.ErrorResponseDto) {
	reqDto.Name = strings.TrimSpace(reqDto.Name)

	if reqDto.Name == "" {
		return nil, model.NewBadRequestErrorDto(util.InvalidRepoNameError)
	}
	req := reqDto.CreateRepoRequest()

	res, err := s.RepoCreator.CreateRepo(req, config.GetGitHubTokenValue())

	if err != nil {
		errDto := err.ErrorResponseDto()

		return nil, errDto
	}
	resDto := res.CreateRepoResponseDto()

	return resDto, nil
}
