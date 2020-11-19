package helper

import (
	"html/template"
	"io"
)

var (
	baseTemplatePath = "templates/base.html"
)

func RenderTemplate(path string, data interface{}, w io.Writer) error {
	templ, err := template.ParseFiles(
		path,
		baseTemplatePath,
	)
	if err != nil {
		return err
	}

	err = templ.ExecuteTemplate(w, "base", data)
	if err != nil {
		return err
	}

	return nil
}
