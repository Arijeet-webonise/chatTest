package app

import (
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/Arijeet-webonise/chatTest/app/constant"
	"github.com/Arijeet-webonise/chatTest/pkg/framework"
	"github.com/Arijeet-webonise/chatTest/pkg/session"
	"golang.org/x/crypto/bcrypt"
)

// RenderUser renders login page
func (app *App) RenderUser(w *framework.Response, r *framework.Request) {
	tplList := []string{
		"web/views/base.html",
		"web/views/user/login.html",
	}

	flash, err := app.SessionManager.GetFlash(w.ResponseWriter, r.Request)
	if err != nil {
		app.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &struct {
		CSRF  template.HTML
		Flash *session.Flash
	}{r.CSRFTemplate(), flash}

	res, err := app.TplParser.ParseTemplate(tplList, data)
	if err != nil {
		app.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w.ResponseWriter, res)
}

// Login process login
func (app *App) Login(w *framework.Response, r *framework.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := app.CustomPortalUserService.PortalUserByEmailID(username)
	if err != nil {
		app.Log.Error(err)
		app.SessionManager.SetFlash(w.ResponseWriter, r.Request, constant.InvalidUserPass, constant.Danger)
		w.Redirect(r.Request, "/login")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		app.Log.Error(err)
		app.SessionManager.SetFlash(w.ResponseWriter, r.Request, constant.InvalidUserPass, constant.Danger)
		w.Redirect(r.Request, "/login")
		return
	}

	sessionID, err := app.SessionManager.SetSession(w.ResponseWriter, r.Request, user.ID)
	if err != nil {
		app.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := app.RedisService.Set(sessionID, user.ID, time.Minute*30); err != nil {
		app.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Logout implement logout
func (app *App) Logout(w *framework.Response, r *framework.Request) {
	userSession, err := app.SessionManager.GetSession(r.Request)
	if err != nil {
		app.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := app.SessionManager.DeleteSession(w.ResponseWriter, r.Request); err != nil {
		app.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := app.RedisService.Delete(userSession.UUID); err != nil {
		app.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Redirect(r.Request, "/")
}
