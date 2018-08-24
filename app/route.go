package app

// InitRoute initcilize routes
func (app *App) InitRoute() {
	app.Router.Get("/", app.renderView(app.RenderIndex))
}
