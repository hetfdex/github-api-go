package service

import (
	"github.com/hetfdex/github-api-go/config"
	"github.com/hetfdex/github-api-go/model"
	"github.com/hetfdex/github-api-go/provider"
	"github.com/hetfdex/github-api-go/util"
	"strings"
	"sync"
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

func (s *service) CreateRepos(reqsDto model.CreateReposRequestDto) (*model.CreateReposResponseDto, *model.ErrorsResponseDto) {
	var wg sync.WaitGroup

	inCh := make(chan createReposChanResult)
	outResCh := make(chan model.CreateReposResponseDto)
	outErrsCh := make(chan model.ErrorsResponseDto)

	defer close(outResCh)
	defer close(outErrsCh)

	go s.handleCreateRepoConcurrentResponse(inCh, outResCh, outErrsCh, &wg)

	for _, reqDto := range reqsDto.Requests {
		wg.Add(1)

		go s.createRepoConcurrent(inCh, reqDto)
	}
	wg.Wait()

	close(inCh)

	responses := <-outResCh
	errors := <-outErrsCh

	return &responses, &errors
}

func (s *service) createRepoConcurrent(inCh chan createReposChanResult, reqDto model.CreateRepoRequestDto) {
	res, err := s.CreateRepo(reqDto)

	resChan := createReposChanResult{
		Response: res,
		Error:    err,
	}
	inCh <- resChan
}

func (s *service) handleCreateRepoConcurrentResponse(inCh chan createReposChanResult, outResCh chan model.CreateReposResponseDto, outErrsCh chan model.ErrorsResponseDto, wg *sync.WaitGroup) {
	var responses model.CreateReposResponseDto
	var errors model.ErrorsResponseDto

	for event := range inCh {
		responses.Responses = append(responses.Responses, *event.Response)
		errors.Errors = append(errors.Errors, *event.Error)

		wg.Done()
	}
	outResCh <- responses
	outErrsCh <- errors
}
