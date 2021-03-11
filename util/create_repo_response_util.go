package util

import (
	"encoding/json"
	"github.com/hetfdex/github-api-go/model"
)

func NewCreateRepoResponseFromBytes(body []byte) (*model.CreateRepoResponse, *model.ErrorResponse) {
	var result model.CreateRepoResponse

	err := json.Unmarshal(body, &result)

	if err != nil {
		return nil, NewInternalServerError(err.Error())
	}
	return &result, nil
}
