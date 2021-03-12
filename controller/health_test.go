package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckHealth(t *testing.T) {
	rec = httptest.NewRecorder()

	ctx, _ = gin.CreateTestContext(rec)

	Health.CheckHealth(ctx)

	assert.EqualValues(t, http.StatusOK, rec.Code)
}
