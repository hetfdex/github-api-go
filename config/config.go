package config

import (
	"github.com/hetfdex/github-api-go/util"
	"os"
)

func GetGitHubTokenValue() string {
	return os.Getenv(util.GitHubTokenKey)
}
