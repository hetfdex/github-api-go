package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const healthy = "healthy"

// CheckHealth endpoint
func CheckHealth(c *gin.Context) {
	c.String(http.StatusOK, healthy)
}
