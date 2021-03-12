package main

import (
	"github.com/hetfdex/github-api-go/router"
)

func main() {
	rtr := router.Router.Setup()

	err := rtr.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}
