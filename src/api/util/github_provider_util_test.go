package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetURL(t *testing.T) {
	url := GetURL("abcde12345")

	assert.NotNil(t, url)
	assert.EqualValues(t, "https://api.github.com/abcde12345", url)
}

func TestGetHeaders(t *testing.T) {
	headers := GetHeaders("abcde12345")

	assert.NotNil(t, headers)
	assert.EqualValues(t, "token abcde12345", headers.Get(autorizationKey))
}

func TestGetAuthorizationValue(t *testing.T) {
	authorizationValue := getAuthorizationValue("abcde12345")

	assert.NotNil(t, authorizationValue)
	assert.EqualValues(t, "token abcde12345", authorizationValue)
}
