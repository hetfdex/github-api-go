package util

import (
	"encoding/json"
	"github.com/hetfdex/github-api-go/dto"
	"github.com/hetfdex/github-api-go/model"
)

func NewCreateRepoResponseFromBytes(body []byte) (*model.CreateRepoResponse, *dto.ErrorResponse) {
	var result model.CreateRepoResponse

	err := json.Unmarshal(body, &result)

	if err != nil {
		return nil, NewInternalServerDtoError(err.Error())
	}
	return &result, nil
}
