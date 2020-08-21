package model

import "net/http"

// APIError model
type APIError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	status  int    `json:"status"`
	message string `json:"message"`
	error   string `json:"error, omitempty"`
}

func (a *apiError) Status() int {
	return a.status
}

func (a *apiError) Message() string {
	return a.message
}

func (a *apiError) Error() string {
	return a.error
}

// NewNotFoundError constructor
func NewNotFoundError(message string) APIError {
	return &apiError{
		status:  http.StatusNotFound,
		message: message,
	}
}

// NewInternalServerError constructor
func NewInternalServerError(message string) APIError {
	return &apiError{
		status:  http.StatusInternalServerError,
		message: message,
	}
}

// NewBadRequestError constructor
func NewBadRequestError(message string) APIError {
	return &apiError{
		status:  http.StatusBadRequest,
		message: message,
	}
}
