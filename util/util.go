package util

import (
	"encoding/json"
	"github.com/hetfdex/github-api-go/model/github"
	"net/http"
)

func NewNotFoundError(message string) *github.ErrorResponse {
	return newError(http.StatusNotFound, message)
}

func NewBadRequestError(message string) *github.ErrorResponse {
	return newError(http.StatusBadRequest, message)
}

func NewInternalServerError(message string) *github.ErrorResponse {
	return newError(http.StatusInternalServerError, message)
}

func newError(statusCode int, message string) *github.ErrorResponse {
	return &github.ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}

func NewErrorFromBytes(statusCode int, body []byte) (*github.ErrorResponse, *github.ErrorResponse) {
	var result github.ErrorResponse

	err := json.Unmarshal(body, &result)

	if err != nil {
		return nil, NewInternalServerError(err.Error())
	}
	result.StatusCode = statusCode

	return &result, nil
}

func NewCreateRepoResponseFromBytes(body []byte) (*github.CreateRepoResponse, *github.ErrorResponse) {
	var result github.CreateRepoResponse

	err := json.Unmarshal(body, &result)

	if err != nil {
		return nil, NewInternalServerError(err.Error())
	}
	return &result, nil
}
