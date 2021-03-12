package mock

import (
	"github.com/gin-gonic/gin"
)

type HealthCheckerMock struct {
	CheckHealthFunc func(ctx *gin.Context)
}

func (*HealthCheckerMock) CheckHealth(*gin.Context) {
	return
}
