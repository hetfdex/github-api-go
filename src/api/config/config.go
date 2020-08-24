package config

import "os"

const gitHubAccessTokenKey = "GITHUB_ACCESSTOKEN_KEY"

var gitHubAccessTokenValue = os.Getenv(gitHubAccessTokenKey)

// GetGitHubAccessToken returns from env
func GetGitHubAccessToken() string {
	return gitHubAccessTokenValue
}
