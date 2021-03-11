package model

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	StatusCode       int     `json:"status_code"`
	Message          string  `json:"message"`
	DocumentationURL string  `json:"documentation_url"`
	Errors           []Error `json:"errors"`
}

type Error struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}

func (e *ErrorResponse) ErrorResponseDto() *ErrorResponseDto {
	return &ErrorResponseDto{
		StatusCode: e.StatusCode,
		Message:    e.Message,
		//Handle Errors
	}
}

func NewInternalServerError(message string) *ErrorResponse {
	return newError(http.StatusInternalServerError, message)
}

func NewErrorFromBytes(statusCode int, body []byte) *ErrorResponse {
	var result ErrorResponse

	err := json.Unmarshal(body, &result)

	if err != nil {
		return newError(http.StatusInternalServerError, err.Error())
	}
	result.StatusCode = statusCode

	return &result
}

func newError(statusCode int, message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}
