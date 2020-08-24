package model

// CreateRepoRequest model
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateRepoResponse model
type CreateRepoResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
