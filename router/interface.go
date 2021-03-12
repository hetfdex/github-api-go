package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hetfdex/github-api-go/controller"
)

type Setuper interface {
	Setup() *gin.Engine
}

type router struct {
	controller.RepoCreator
	controller.HealthChecker
}
