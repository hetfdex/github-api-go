package config

import "os"

const gitHubTokenKey = "SECRET_GITHUB_TOKEN"

var gitHubTokenValue = os.Getenv(gitHubTokenKey)

func GetGitHubTokenValue() string {
	return gitHubTokenValue
}
