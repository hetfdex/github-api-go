package model

type CreateRepoResponseDto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type CreateReposResponseDto struct {
	StatusCode int                     `json:"status_code"`
	Responses  []CreateRepoResponseDto `json:"responses"`
}

type CreateReposResponseChan struct {
	Response CreateRepoResponseDto `json:"response"`
	Error    ErrorResponseDto      `json:"error"`
}

func NewCreateReposResponseChan(res CreateRepoResponse, err ErrorResponse) *CreateReposResponseChan {
	return &CreateReposResponseChan{
		Response: *res.CreateRepoResponseDto(),
		Error:    *err.ErrorResponseDto(),
	}
}
