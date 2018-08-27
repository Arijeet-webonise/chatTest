package app

import (
	"io"
	"net/http"

	"github.com/Arijeet-webonise/chatTest/pkg/framework"
	"github.com/Arijeet-webonise/chatTest/pkg/session"
)

// RenderIndex renders index page
func (app *App) RenderIndex(w *framework.Response, r *framework.Request) {
	tplList := []string{
		"web/views/base.html",
		"web/views/index.html",
	}
	flash, err := app.SessionManager.GetFlash(w.ResponseWriter, r.Request)
	if err != nil {
		app.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &struct {
		Flash *session.Flash
	}{flash}

	res, err := app.TplParser.ParseTemplate(tplList, data)
	if err != nil {
		app.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w.ResponseWriter, res)
}
