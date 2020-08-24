package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

// StartApp starts the webserver
func StartApp() {
	mapURLs()

	err := router.Run()

	if err != nil {
		panic(err)
	}
}