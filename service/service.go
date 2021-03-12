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
	if !isValidName(reqDto.Name) {
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

func (s *service) CreateRepos(reqs model.CreateReposRequestDto) (*model.CreateReposResponseDto, *model.ErrorResponseDto) {
	c := make(chan *model.CreateReposResponseDto)

	for _, repo := range reqs.Requests {
		go createRepoConcurrent(repo, c)
	}
	return nil, nil
}

func createRepoConcurrent(req model.CreateRepoRequestDto, ch chan *model.CreateReposResponseDto) {
	if !isValidName(req.Name) {
		ch <- nil
	}
	//ch <-
}

func isValidName(name string) bool {
	name = strings.TrimSpace(name)

	if name == "" {
		return false
	}
	return true
}
