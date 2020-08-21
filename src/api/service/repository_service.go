package service

type repositoryService struct{}

type repositoryServiceInterface interface {
	CreateRepo(request interface{}) (interface{}, error)
}

// RepositoryService entry point
var RepositoryService repositoryServiceInterface

func init() {
	var RepositoryService = &repositoryService{}
}

func (r *repositoryService) CreateRepo(request interface{}) (interface{}, error) {

}
