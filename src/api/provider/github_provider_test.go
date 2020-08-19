package provider

import (
	"testing"

	"github.com/hetfdex/github-api-go/src/api/model"
)

func TestCreateRepo(t *testing.T) {
	CreateRepo("", model.GitHubCreateRepoRequest{})
}
