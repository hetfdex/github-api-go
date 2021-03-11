package util

import (
	"encoding/json"
	"github.com/hetfdex/github-api-go/model"
	"net/http"
)

func NewNotFoundError(message string) *model.ErrorResponse {
	return newError(http.StatusNotFound, message)
}

func NewBadRequestError(message string) *model.ErrorResponse {
	return newError(http.StatusBadRequest, message)
}

func NewInternalServerError(message string) *model.ErrorResponse {
	return newError(http.StatusInternalServerError, message)
}

func NewErrorFromBytes(statusCode int, body []byte) (*model.ErrorResponse, *model.ErrorResponse) {
	var result model.ErrorResponse

	err := json.Unmarshal(body, &result)

	if err != nil {
		return nil, NewInternalServerError(err.Error())
	}
	result.StatusCode = statusCode

	return &result, nil
}

func newError(statusCode int, message string) *model.ErrorResponse {
	return &model.ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}
