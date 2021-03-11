package model

type CreateRepoResponseDto struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Owner       string      `json:"owner"`
}
