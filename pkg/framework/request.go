package framework

import (
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

// Request framework for Request
type Request struct {
	*http.Request
	context map[string]interface{}
}

//CSRFTemplate generate HTML template for CSRF
func (res *Request) CSRFTemplate() template.HTML {
	return csrf.TemplateField(res.Request)
}

// Push pushed value to a map
func (res *Request) Push(key string, value interface{}) {
	if res.context == nil {
		res.context = make(map[string]interface{})
	}
	res.context[key] = value
}

// Value return value from map
func (res *Request) Value(key string) interface{} {
	value := res.context[key]
	return value
}
