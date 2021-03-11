package service

import (
	"github.com/hetfdex/github-api-go/provider"
)

type serviceInterface interface {
}

type service struct {
	provider.RepoCreator
}

var Service serviceInterface = &service{
	provider.Provider,
}
