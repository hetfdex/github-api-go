package app

func mapRoutes(a *app) {
	engine.GET("/health-check", a.CheckHealth)
	engine.POST("/repositories", a.CreateRepo)
}
