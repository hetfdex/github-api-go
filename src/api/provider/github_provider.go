package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hetfdex/github-api-go/src/api/client"
	"github.com/hetfdex/github-api-go/src/api/model/github"
)

const baseURL = "https://api.github.com/%s"
const createRepoPath = "user/repos"
const autorizationKey = "Authorization"
const autorizationValue = "token %s"

func getURL(path string) string {
	return fmt.Sprintf(baseURL, path)
}

func getHeaders(authorizationToken string) http.Header {
	authorizationValue := getAuthorizationValue(authorizationToken)

	headers := http.Header{}

	headers.Set(autorizationKey, authorizationValue)

	return headers
}

func getAuthorizationValue(authorizationToken string) string {
	return fmt.Sprintf(autorizationValue, authorizationToken)
}

// CreateRepo requests the creation of a new github repository
func CreateRepo(authorizationToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	url := getURL(createRepoPath)

	headers := getHeaders(authorizationToken)

	response, err := client.Post(url, request, headers)

	if err != nil {
		log.Println(fmt.Sprintf("Github CreateRepo Request Error: %s", err.Error()))

		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid Error Response Body",
		}
	}

	if response.StatusCode > 299 {
		var errResponse github.ErrorResponse

		err := json.Unmarshal(bytes, &errResponse)

		if err != nil {
			return nil, &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid JSON Error Response Body",
			}
		}
		errResponse.StatusCode = response.StatusCode

		return nil, &errResponse
	}
	var result github.CreateRepoResponse

	err = json.Unmarshal(bytes, &result)

	if err != nil {
		log.Println(fmt.Sprintf("Github CreateRepo Response Error: %s", err.Error()))

		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid JSON Success Response Body",
		}
	}
	return &result, nil
}
