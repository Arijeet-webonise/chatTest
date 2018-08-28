package templates

import (
	"bytes"
	"html/template"
)

// TplParse encapsulate template
type TplParse interface {
	ParseTemplate(templateFileName []string, data interface{}) (string, error)
}

// TemplateParser implement TplParse
type TemplateParser struct {
}

//ParseTemplate convert template and related data into html formated text
func (tp *TemplateParser) ParseTemplate(templateFileName []string, data interface{}) (string, error) {
	var parsedTemplate string

	//always execute flash template
	templateFileName = append(templateFileName, "./web/views/layouts/flash_message.html")

	t, err := template.ParseFiles(templateFileName...)
	if err != nil {
		return parsedTemplate, err
	}
	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		return parsedTemplate, err
	}
	parsedTemplate = buf.String()

	return parsedTemplate, nil
}
