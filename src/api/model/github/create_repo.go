package github

// CreateRepoRequest model
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

// CreateRepoResponse model
type CreateRepoResponse struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	FullName    string      `json:"full_name"`
	Owner       Owner       `json:"owner"`
	Permissions Permissions `json:"permissions"`
}

// Owner model
type Owner struct {
	Login   string `json:"login"`
	ID      int    `json:"id"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

//Permissions model
type Permissions struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}
