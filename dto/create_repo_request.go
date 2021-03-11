package dto

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
