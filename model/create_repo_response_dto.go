package model

type CreateRepoResponseDto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type CreateReposResponseDto struct {
	Responses []CreateRepoResponseDto `json:"responses"`
	Error     ErrorResponseDto        `json:"error"`
}
