package mapper

import (
	"fmt"
	"github.com/hetfdex/github-api-go/dto"
	"github.com/hetfdex/github-api-go/model"
)

func ToErrorResponseDto(res *model.ErrorResponse) *dto.ErrorResponse {
	return &dto.ErrorResponse{
		StatusCode: res.StatusCode,
		Message:    res.Message,
		Error:      toErrorDto(res.Errors),
	}
}

func toErrorDto(errors []model.Error) string {
	var result string

	for _, err := range errors {
		result = result + fmt.Sprintf("%s", err.Message)
	}
	return result
}
