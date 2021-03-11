package model

type CreateRepoRequestDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c CreateRepoRequestDto) CreateRepoRequest() CreateRepoRequest {
	return CreateRepoRequest{
		Name:        c.Name,
		Description: c.Description,
		Private:     false,
	}
}
