package model

import "net/http"

// APIError model
type APIError interface {
	StatusCode() int
	Message() string
	Error() string
}

type apiError struct {
	statusCode int    `json:"status_code"`
	message    string `json:"message"`
	error      string `json:"error,omitempty"`
}

func (a *apiError) StatusCode() int {
	return a.statusCode
}

func (a *apiError) Message() string {
	return a.message
}

func (a *apiError) Error() string {
	return a.error
}

// NewAPIError constructor
func NewAPIError(statusCode int, message string) APIError {
	return &apiError{
		statusCode: statusCode,
		message:    message,
	}
}

// NewNotFoundError constructor
func NewNotFoundError(message string) APIError {
	return &apiError{
		statusCode: http.StatusNotFound,
		message:    message,
	}
}

// NewInternalServerError constructor
func NewInternalServerError(message string) APIError {
	return &apiError{
		statusCode: http.StatusInternalServerError,
		message:    message,
	}
}

// NewBadRequestError constructor
func NewBadRequestError(message string) APIError {
	return &apiError{
		statusCode: http.StatusBadRequest,
		message:    message,
	}
}
