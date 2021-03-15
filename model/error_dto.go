package model

import (
	"net/http"
)

type ErrorResponseDto struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewInternalServerErrorDto(message string) *ErrorResponseDto {
	return newErrorDto(http.StatusInternalServerError, message)
}

func NewBadRequestErrorDto(message string) *ErrorResponseDto {
	return newErrorDto(http.StatusBadRequest, message)
}

func newErrorDto(statusCode int, message string) *ErrorResponseDto {
	return &ErrorResponseDto{
		StatusCode: statusCode,
		Message:    message,
	}
}
