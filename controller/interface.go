package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hetfdex/github-api-go/service"
)

type RepoCreator interface {
	CreateRepo(ctx *gin.Context)
	CreateRepos(ctx *gin.Context)
}

type HealthChecker interface {
	CheckHealth(ctx *gin.Context)
}

type controller struct {
	service.RepoCreator
}

type health struct {
}
