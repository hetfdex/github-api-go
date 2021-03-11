package client

import (
	"bytes"
	"github.com/hetfdex/github-api-go/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

var reader = bytes.NewReader([]byte(`test`))

var header = http.Header{}

func TestMain(m *testing.M) {
	httpClient = &mock.DoerMock{}

	os.Exit(m.Run())
}

func TestPostNewRequestError(t *testing.T) {
	url := ":abc1{DEf2=test@test.com:666/db?"

	res, err := PostClient.Post(url, header, reader)

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

func TestPostError(t *testing.T) {
	url := mock.DoErrorUrl

	res, err := PostClient.Post(url, header, reader)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, mock.DoErrorUrl, err.Error())
}

func TestPostOk(t *testing.T) {
	url := "https://www.google.com"

	res, err := PostClient.Post(url, header, reader)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
}
