package util

import (
	"encoding/json"
	"github.com/hetfdex/github-api-go/model"
	"net/http"
)

func NewInternalServerError(message string) *model.ErrorResponse {
	return newError(http.StatusInternalServerError, message)
}

func NewErrorFromBytes(statusCode int, body []byte) *model.ErrorResponse {
	var result model.ErrorResponse

	err := json.Unmarshal(body, &result)

	if err != nil {
		return newError(http.StatusInternalServerError, err.Error())
	}
	result.StatusCode = statusCode

	return &result
}

func newError(statusCode int, message string) *model.ErrorResponse {
	return &model.ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}
