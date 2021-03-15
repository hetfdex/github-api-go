package model

import "encoding/json"

type CreateRepoResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c CreateRepoResponse) CreateRepoResponseDto() *CreateRepoResponseDto {
	return &CreateRepoResponseDto{
		ID:   c.ID,
		Name: c.Name,
	}
}

func NewCreateRepoResponseFromBytes(body []byte) (*CreateRepoResponse, *ErrorResponse) {
	var result CreateRepoResponse

	err := json.Unmarshal(body, &result)

	if err != nil {
		return nil, NewInternalServerError(err.Error())
	}
	return &result, nil
}
