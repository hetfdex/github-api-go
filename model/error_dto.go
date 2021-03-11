package model

import (
	"net/http"
)

type ErrorResponseDto struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error,omitempty"`
}

func NewInternalServerDtoError(message string) *ErrorResponseDto {
	return newDtoError(http.StatusInternalServerError, message)
}

func NewBadRequestDtoError(message string) *ErrorResponseDto {
	return newDtoError(http.StatusBadRequest, message)
}

func newDtoError(statusCode int, message string) *ErrorResponseDto {
	return &ErrorResponseDto{
		StatusCode: statusCode,
		Message:    message,
	}
}
