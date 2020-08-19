package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hetfdex/github-api-go/src/api/client"
	"github.com/hetfdex/github-api-go/src/api/model"
	"github.com/hetfdex/github-api-go/src/api/util"
)

// CreateRepo requests the creation of a new github repository
func CreateRepo(authorizationToken string, request model.GitHubCreateRepoRequest) (*model.GitHubCreateRepoResponse, *model.GitHubErrorResponse) {
	url := util.GetURL(util.CreateRepoPath)

	headers := util.GetHeaders(authorizationToken)

	response, err := client.Post(url, request, headers)

	if err != nil {
		log.Println(fmt.Sprintf("Github CreateRepo Request Error: %s", err.Error()))

		return nil, &model.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, &model.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid Error Response Body",
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse model.GitHubErrorResponse

		err := json.Unmarshal(bytes, &errResponse)

		if err != nil {
			return nil, &model.GitHubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid JSON Error Response Body",
			}
		}
		errResponse.StatusCode = response.StatusCode

		return nil, &errResponse
	}
	var result model.GitHubCreateRepoResponse

	err = json.Unmarshal(bytes, &result)

	if err != nil {
		log.Println(fmt.Sprintf("Github CreateRepo Response Error: %s", err.Error()))

		return nil, &model.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid JSON Success Response Body",
		}
	}
	return &result, nil
}
