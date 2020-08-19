package util

import (
	"fmt"
	"net/http"
)

const baseURL = "https://api.github.com/%s"
const autorizationKey = "Authorization"
const autorizationValue = "token %s"

// CreateRepoPath for url
const CreateRepoPath = "user/repos"

// GetURL appends path to base url
func GetURL(path string) string {
	return fmt.Sprintf(baseURL, path)
}

// GetHeaders returns headers with authorization
func GetHeaders(authorizationToken string) http.Header {
	authorizationValue := getAuthorizationValue(authorizationToken)

	headers := http.Header{}

	headers.Set(autorizationKey, authorizationValue)

	return headers
}

func getAuthorizationValue(authorizationToken string) string {
	return fmt.Sprintf(autorizationValue, authorizationToken)
}
