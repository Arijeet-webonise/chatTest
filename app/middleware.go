package app

import (
	"net/http"

	"github.com/Arijeet-webonise/chatTest/pkg/framework"
)

func (app *App) renderView(viewHandler func(*framework.Response, *framework.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := framework.Response{ResponseWriter: w}
		req := framework.Request{Request: r}
		viewHandler(&res, &req)
	})
}
