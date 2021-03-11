package model

import "encoding/json"

type CreateRepoResponse struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	FullName    string      `json:"full_name"`
	Owner       Owner       `json:"owner"`
	Permissions Permissions `json:"permissions"`
}

type Owner struct {
	Login   string `json:"login"`
	ID      int    `json:"id"`
	URL     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

type Permissions struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}

func NewCreateRepoResponseFromBytes(body []byte) (*CreateRepoResponse, *ErrorResponse) {
	var result CreateRepoResponse

	err := json.Unmarshal(body, &result)

	if err != nil {
		return nil, NewInternalServerError(err.Error())
	}
	return &result, nil
}
