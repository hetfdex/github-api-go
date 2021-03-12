package app

import (
	"github.com/hetfdex/github-api-go/controller"
)

type Starter interface {
	Start()
}

type app struct {
	controller.RepoCreator
}
