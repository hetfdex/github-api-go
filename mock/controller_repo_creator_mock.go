package mock

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const createRepoCreated = "createRepoCreated"

type ControllerRepoCreatorMock struct {
	CreateRepoFunc func(ctx *gin.Context)
}

func (*ControllerRepoCreatorMock) CreateRepo(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, createRepoCreated)

	return
}
