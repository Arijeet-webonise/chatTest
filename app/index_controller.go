package app

import (
	"io"
	"net/http"
)

// RenderIndex renders index page
func (app *App) RenderIndex(w http.ResponseWriter, r *http.Request) {
	tplList := []string{
		"web/views/base.html",
	}

	res, err := app.TplParser.ParseTemplate(tplList, nil)
	if err != nil {
		app.Log.Error("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, res)
}
