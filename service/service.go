package service

import (
	"github.com/hetfdex/github-api-go/config"
	"github.com/hetfdex/github-api-go/model"
	"github.com/hetfdex/github-api-go/provider"
	"github.com/hetfdex/github-api-go/util"
	"net/http"
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

func (s *service) CreateRepos(reqsDto model.CreateReposRequestDto) *model.CreateReposResponseDto {
	var wg sync.WaitGroup

	inCh := make(chan createReposChanResult)
	outCh := make(chan model.CreateReposResponseDto)

	defer close(outCh)

	go s.handleCreateRepoConcurrentResponse(inCh, outCh, &wg)

	for _, reqDto := range reqsDto.Requests {
		wg.Add(1)

		go s.createRepoConcurrent(inCh, reqDto)
	}
	wg.Wait()

	close(inCh)

	responses := <-outCh

	responses.StatusCode = createReposStatusCode(responses, len(reqsDto.Requests))

	return &responses
}

func (s *service) createRepoConcurrent(inCh chan createReposChanResult, reqDto model.CreateRepoRequestDto) {
	res, err := s.CreateRepo(reqDto)

	resChan := createReposChanResult{
		Response: res,
		Error:    err,
	}
	inCh <- resChan
}

func (s *service) handleCreateRepoConcurrentResponse(inCh chan createReposChanResult, outCh chan model.CreateReposResponseDto, wg *sync.WaitGroup) {
	var responses model.CreateReposResponseDto

	for event := range inCh {
		responses.Responses = append(responses.Responses, *event.Response)
		responses.Errors = append(responses.Errors, *event.Error)

		wg.Done()
	}
	outCh <- responses
}

func createReposStatusCode(ress model.CreateReposResponseDto, reqCount int) int {
	successCount := 0
	failureCount := 0

	for _, res := range ress.Responses {
		if res.ID != 0 && res.Name != "" {
			successCount++
		}
	}

	for _, err := range ress.Errors {
		if err.StatusCode != 0 && err.Message != "" {
			failureCount++
		}
	}

	if successCount == reqCount {
		return http.StatusCreated
	}

	if failureCount == reqCount {
		return http.StatusInternalServerError
	}
	return http.StatusPartialContent
}
