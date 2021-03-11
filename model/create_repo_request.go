package model

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoRequestDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c CreateRepoRequestDto) ToCreateRepoRequest() CreateRepoRequest {
	return CreateRepoRequest{
		Name:        c.Name,
		Description: c.Description,
		Private:     false,
	}
}
