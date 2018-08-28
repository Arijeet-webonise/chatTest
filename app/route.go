package app

// InitRoute initcilize routes
func (app *App) InitRoute() {
	app.Router.Get("/", app.renderView(app.RenderIndex))
	app.Router.Get("/login", app.renderView(app.RenderUser))
	app.Router.Get("/logout", app.renderView(app.Logout))
	app.Router.Post("/login", app.renderView(app.Login))
}
