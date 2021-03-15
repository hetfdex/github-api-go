package model

type CreateRepoResponseDto struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateReposResponseDto struct {
	StatusCode int                      `json:"status_code"`
	Responses  []*CreateRepoResponseDto `json:"responses"`
	Errors     []*ErrorResponseDto      `json:"errors"`
}
