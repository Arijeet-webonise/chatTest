package app

import (
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Arijeet-webonise/chatTest/app/constant"
	"github.com/Arijeet-webonise/chatTest/app/models"
	"github.com/Arijeet-webonise/chatTest/pkg/framework"
	"github.com/Arijeet-webonise/chatTest/pkg/session"
	"github.com/go-zoo/bone"
	"golang.org/x/crypto/bcrypt"
)

// RenderUser renders login page
func (app *App) RenderUser(w *framework.Response, r *framework.Request) {
	tplList := []string{
		"web/views/base.html",
		"web/views/menu/NONE.html",
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

	w.Redirect(r.Request, "/profile")
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

// GetCurrentUser returns current user
func (app *App) GetCurrentUser(w *framework.Response, r *framework.Request) (*models.PortalUser, error) {
	userSession, err := app.SessionManager.GetSession(r.Request)
	if err != nil {
		app.Log.Error(err)
		return nil, err
	}
	userID := 0
	result, err := app.RedisService.Get(userSession.UUID)
	if err != nil {
		app.Log.Error(err)
		return nil, err
	}
	if err := result.Scan(&userID); err != nil {
		app.Log.Error(err)
		return nil, err
	}

	user, err := app.CustomPortalUserService.PortalUserByID(userID)
	if err != nil {
		app.Log.Error(err)
		return nil, err
	}
	return user, nil
}

func (app *App) RenderCurrentUser(w *framework.Response, r *framework.Request) {
	user := r.Value("user").(*models.PortalUser)
	tplList := []string{
		"web/views/base.html",
		"web/views/menu/NONE.html",
		"web/views/user/profile.html",
	}

	flash, err := app.SessionManager.GetFlash(w.ResponseWriter, r.Request)
	if err != nil {
		app.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &struct {
		User  *models.PortalUser
		Flash *session.Flash
	}{user, flash}

	res, err := app.TplParser.ParseTemplate(tplList, data)

	w.RenderHTML(res)
}

func (app *App) RenderUserProfile(w *framework.Response, r *framework.Request) {
	tplList := []string{
		"web/views/base.html",
		"web/views/menu/NONE.html",
		"web/views/user/profile.html",
	}

	userID, err := strconv.Atoi(bone.GetValue(r.Request, "id"))
	if err != nil {
		http.Error(w.ResponseWriter, "Invalid ID", http.StatusInternalServerError)
		return
	}

	user, err := app.CustomPortalUserService.PortalUserByID(userID)
	if err != nil {
		app.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	flash, err := app.SessionManager.GetFlash(w.ResponseWriter, r.Request)
	if err != nil {
		app.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &struct {
		User  *models.PortalUser
		Flash *session.Flash
	}{user, flash}

	res, err := app.TplParser.ParseTemplate(tplList, data)

	w.RenderHTML(res)
}
