package model

// GitHubCreateRepoRequest model
type GitHubCreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

// GitHubCreateRepoResponse model
type GitHubCreateRepoResponse struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	FullName    string            `json:"full_name"`
	Owner       GitHubOwner       `json:"owner"`
	Permissions GitHubPermissions `json:"permissions"`
}

// GitHubOwner model
type GitHubOwner struct {
	Login   string `json:"login"`
	ID      int    `json:"id"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

// GitHubPermissions model
type GitHubPermissions struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}
