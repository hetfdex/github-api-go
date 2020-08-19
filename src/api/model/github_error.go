package model

// GitHubErrorResponse model
type GitHubErrorResponse struct {
	StatusCode       int           `json:"status_code"`
	Message          string        `json:"message"`
	DocumentationURL string        `json:"documentation_url"`
	Errors           []GitHubError `json:"errors"`
}

// GitHubError model
type GitHubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
