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

func (s *service) CreateRepos(requestsDto model.CreateReposRequestDto) *model.CreateReposResponseDto {
	var wg sync.WaitGroup

	inCh := make(chan createReposChanResult)
	outCh := make(chan model.CreateReposResponseDto)
	buffCh := make(chan int, util.ConcurrencyBuffer)

	defer close(outCh)

	go s.handleCreateRepoConcurrentResponse(inCh, outCh, &wg)

	for i, reqDto := range requestsDto.Requests {
		buffCh <- i

		wg.Add(1)

		go s.createRepoConcurrent(inCh, buffCh, reqDto)
	}
	wg.Wait()

	close(inCh)

	responsesDto := <-outCh

	responsesDto.StatusCode = getStatusCode(responsesDto, len(requestsDto.Requests))

	return &responsesDto
}

func (s *service) createRepoConcurrent(inCh chan createReposChanResult, buffCh chan int, reqDto model.CreateRepoRequestDto) {
	res, err := s.CreateRepo(reqDto)

	resChan := createReposChanResult{
		Response: res,
		Error:    err,
	}
	inCh <- resChan

	<-buffCh
}

func (s *service) handleCreateRepoConcurrentResponse(inCh chan createReposChanResult, outCh chan model.CreateReposResponseDto, wg *sync.WaitGroup) {
	var responsesDto model.CreateReposResponseDto

	for event := range inCh {
		if event.Response != nil {
			responsesDto.Responses = append(responsesDto.Responses, *event.Response)
		}

		if event.Error != nil {
			responsesDto.Errors = append(responsesDto.Errors, *event.Error)
		}
		wg.Done()
	}
	outCh <- responsesDto
}

func getStatusCode(responsesDto model.CreateReposResponseDto, reqCount int) int {
	successCount := 0
	failureCount := 0

	for _, resDto := range responsesDto.Responses {
		if resDto.ID != 0 && resDto.Name != "" {
			successCount++
		}
	}

	for _, errDto := range responsesDto.Errors {
		if errDto.StatusCode != 0 && errDto.Message != "" {
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
