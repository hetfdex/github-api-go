package service

import "github.com/hetfdex/github-api-go/provider/github"

type serviceInterface interface {
}

type service struct {
	github.RepoCreator
}

var Service serviceInterface = &service{
	github.Provider,
}
