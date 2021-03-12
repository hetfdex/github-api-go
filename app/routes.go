package app

func mapRoutes(a *app) {
	engine.POST("/repositories", a.CreateRepo)
}
