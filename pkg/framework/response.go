package framework

import (
	"io"
	"net/http"
)

// Response framework for Response
type Response struct {
	http.ResponseWriter
}

// Redirect redirects to a url
func (res *Response) Redirect(r *http.Request, url string) {
	http.Redirect(res.ResponseWriter, r, url, http.StatusSeeOther)
}

// RenderHTML renders HTML pages
func (res *Response) RenderHTML(htmlStr string) {
	io.WriteString(res.ResponseWriter, htmlStr)
}
