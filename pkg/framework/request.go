package framework

import (
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

// Request framework for Request
type Request struct {
	*http.Request
}

//CSRFTemplate generate HTML template for CSRF
func (res *Request) CSRFTemplate() template.HTML {
	return csrf.TemplateField(res.Request)
}
