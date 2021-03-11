package model

type ErrorResponseDto struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error     string `json:"error, omitempty"`
}
