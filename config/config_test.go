package config

import (
	"github.com/hetfdex/github-api-go/util"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetGitHubTokenValue(t *testing.T) {
	mockValue := "mockValue"

	_ = os.Setenv(util.GitHubTokenKey, mockValue)

	result := GetGitHubTokenValue()

	assert.EqualValues(t, mockValue, result)

	t.Cleanup(
		func() {
			os.Clearenv()
		},
	)
}
