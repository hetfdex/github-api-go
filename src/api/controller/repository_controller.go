package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hetfdex/github-api-go/src/api/model"
	"github.com/hetfdex/github-api-go/src/api/service"
)

// CreateRepo endpoint
func CreateRepo(c *gin.Context) {
	var request model.CreateRepoRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		error := model.NewBadRequestError("Invalid JSON Body")

		c.JSON(error.StatusCode(), error)

		return
	}
	response, err := service.RepositoryService.CreateRepo(request)

	if err != nil {
		c.JSON(err.StatusCode(), err)

		return
	}
	c.JSON(http.StatusCreated, response)
}
