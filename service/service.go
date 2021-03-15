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
	name := strings.TrimSpace(reqDto.Name)

	if name == "" {
		return nil, model.NewBadRequestErrorDto(util.InvalidRepoNameError)
	}
	req := reqDto.CreateRepoRequest()

	res, err := s.RepoCreator.CreateRepo(*req, config.GetGitHubTokenValue())

	if err != nil {
		errDto := err.ErrorResponseDto()

		return nil, errDto
	}
	resDto := res.CreateRepoResponseDto()

	return resDto, nil
}

func (s *service) CreateRepos(reqsDto model.CreateReposRequestDto) (*model.CreateReposResponseDto, *model.ErrorResponseDto) {
	inCh := make(chan createReposChanReturn)
	outCh := make(chan model.CreateReposResponseDto)

	go s.handleCreateRepoConcurrentResponse(inCh, outCh)

	for _, reqDto := range reqsDto.Requests {
		go s.createRepoConcurrent(inCh, reqDto)
	}
	responses := <-outCh

	return &responses, nil
}

func (s *service) createRepoConcurrent(inCh chan createReposChanReturn, reqDto model.CreateRepoRequestDto) {
	res, err := s.CreateRepo(reqDto)

	resChan := createReposChanReturn{
		Response: res,
		Error:    err,
	}
	inCh <- resChan
}

func (s *service) handleCreateRepoConcurrentResponse(inCh chan createReposChanReturn, outCh chan model.CreateReposResponseDto) {
	var responses model.CreateReposResponseDto

	for response := range inCh {
		responses.Responses = append(responses.Responses, *response.Response)
	}
	outCh <- responses
}
