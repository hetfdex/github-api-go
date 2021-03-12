package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hetfdex/github-api-go/model"
	"github.com/hetfdex/github-api-go/service"
	"github.com/hetfdex/github-api-go/util"
	"net/http"
)

var Controller RepoCreator = &controller{
	service.Service,
}

func (c *controller) CreateRepo(ctx *gin.Context) {
	var req model.CreateRepoRequestDto

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		er := model.NewBadRequestErrorDto(util.InvalidJsonBodyError)

		ctx.JSON(er.StatusCode, er)

		return
	}
	res, er := c.RepoCreator.CreateRepo(req)

	if er != nil {
		ctx.JSON(er.StatusCode, er)

		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (c *controller) CreateRepos(ctx *gin.Context) {
	var req model.CreateReposRequestDto

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		er := model.NewBadRequestErrorDto(util.InvalidJsonBodyError)

		ctx.JSON(er.StatusCode, er)

		return
	}
	res, er := c.RepoCreator.CreateRepos(req)

	if er != nil {
		ctx.JSON(er.StatusCode, er)

		return
	}
	ctx.JSON(http.StatusCreated, res)
}
