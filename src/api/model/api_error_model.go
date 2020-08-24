package model

import "net/http"

// APIError model
type APIError interface {
	GetStatusCode() int
	GetMessage() string
	GetError() string
}

type apiError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error,omitempty"`
}

func (a *apiError) GetStatusCode() int {
	return a.StatusCode
}

func (a *apiError) GetMessage() string {
	return a.Message
}

func (a *apiError) GetError() string {
	return a.Error
}

// NewAPIError constructor
func NewAPIError(statusCode int, message string) APIError {
	return &apiError{
		StatusCode: statusCode,
		Message:    message,
	}
}

// NewNotFoundError constructor
func NewNotFoundError(message string) APIError {
	return &apiError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

// NewInternalServerError constructor
func NewInternalServerError(message string) APIError {
	return &apiError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

// NewBadRequestError constructor
func NewBadRequestError(message string) APIError {
	return &apiError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}
