package app

import "net/http"

// InitRoute initcilize routes
func (app *App) InitRoute() {
	app.Router.Get("/", app.renderView(app.RenderIndex))
	app.Router.Get("/login", app.renderView(app.RenderUser))
	app.Router.Post("/login", app.renderView(app.Login))
	app.Router.Get("/logout", app.renderView(app.Logout))
	app.Router.Get("/profile", app.renderSecureView(app.RenderCurrentUser))
	app.Router.Get("/profile/:id", app.renderSecureView(app.RenderCurrentUser))

	fs := http.FileServer(http.Dir("web/assets"))
	app.Router.Get("/static/", http.StripPrefix("/static/", fs))
}
