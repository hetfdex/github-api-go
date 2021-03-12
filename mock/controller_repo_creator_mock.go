package mock

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerRepoCreatorMock struct {
	CreateRepoFunc  func(ctx *gin.Context)
	CreateReposFunc func(ctx *gin.Context)
}

func (*ControllerRepoCreatorMock) CreateRepo(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, http.StatusAccepted)
}

func (*ControllerRepoCreatorMock) CreateRepos(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, http.StatusAccepted)
}
