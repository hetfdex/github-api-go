package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Health HealthChecker = &health{}

func (h *health) CheckHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, http.StatusOK)
}
