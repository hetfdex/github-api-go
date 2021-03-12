package mock

import (
	"github.com/gin-gonic/gin"
)

type ControllerRepoCreatorMock struct {
	CreateRepoFunc func(ctx *gin.Context)
}

func (*ControllerRepoCreatorMock) CreateRepo(ctx *gin.Context) {
	return
}
