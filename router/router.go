package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hetfdex/github-api-go/controller"
	"github.com/hetfdex/github-api-go/util"
)

var Router Setuper = &router{
	controller.Controller,
	controller.Health,
}

func (e *router) Setup() *gin.Engine {
	engine := gin.Default()

	engine.GET(util.HealthCheckPath, e.CheckHealth)
	engine.POST(util.RepositoryPath, e.CreateRepo)
	engine.POST(util.RepositoriesPath, e.CreateRepos)

	return engine
}
