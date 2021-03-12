package mock

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckerMock struct {
	CheckHealthFunc func(ctx *gin.Context)
}

func (*HealthCheckerMock) CheckHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, http.StatusAccepted)
}
