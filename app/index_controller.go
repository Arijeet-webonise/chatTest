package app

import (
	"io"
	"log"
	"net/http"
)

// RenderIndex renders index page
func (app *App) RenderIndex(w http.ResponseWriter, r *http.Request) {
	tplList := []string{
		"web/views/base.html",
	}

	res, err := app.TplParser.ParseTemplate(tplList, nil)
	if err != nil {
		log.Panicln(err)
	}

	io.WriteString(w, res)
}
