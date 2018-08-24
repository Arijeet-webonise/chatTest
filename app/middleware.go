package app

import "net/http"

func (app *App) renderView(viewHandler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		viewHandler(w, r)
	})
}
