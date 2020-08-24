package app

import "github.com/hetfdex/github-api-go/src/api/controller"

func mapURLs() {
	router.GET("/health-check", controller.CheckHealth)
	router.POST("/repositories", controller.CreateRepo)
}
