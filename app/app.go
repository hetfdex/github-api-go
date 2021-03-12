package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hetfdex/github-api-go/controller"
)

var App Starter = &app{
	controller.Controller,
	controller.Health,
}

var engine = gin.Default()

func (a *app) Start() {
	mapRoutes(a)

	err := engine.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}
