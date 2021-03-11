package util

import (
	"encoding/json"
	"github.com/hetfdex/github-api-go/dto"
	"net/http"
)

func NewNotFoundDtoError(message string) *dto.ErrorResponse {
	return newDtoError(http.StatusNotFound, message)
}

func NewBadRequestDtoError(message string) *dto.ErrorResponse {
	return newDtoError(http.StatusBadRequest, message)
}

func NewInternalServerDtoError(message string) *dto.ErrorResponse {
	return newDtoError(http.StatusInternalServerError, message)
}

func NewDtoErrorFromBytes(statusCode int, body []byte) *dto.ErrorResponse {
	var result dto.ErrorResponse

	err := json.Unmarshal(body, &result)

	if err != nil {
		return newDtoError(http.StatusInternalServerError, err.Error())
	}
	result.StatusCode = statusCode

	return &result
}

func newDtoError(statusCode int, message string) *dto.ErrorResponse {
	return &dto.ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}
