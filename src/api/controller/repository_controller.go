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

	error := c.ShouldBindJSON(&request)

	if error != nil {
		err := model.NewBadRequestError("Invalid JSON Body")

		c.JSON(err.GetStatusCode(), err)

		return
	}
	response, err := service.RepositoryService.CreateRepo(request)

	if err != nil {
		c.JSON(err.GetStatusCode(), err)

		return
	}
	c.JSON(http.StatusCreated, response)
}
