package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestCheckHealth(t *testing.T) {
	res := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(res)

	request, _ := http.NewRequest(http.MethodPost, "/url", strings.NewReader(``))

	c.Request = request

	CheckHealth(c)

	assert.EqualValues(t, http.StatusOK, res.Code)
}
