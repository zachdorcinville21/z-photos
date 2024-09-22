package main

import (
	"io"

	"html/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	app := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob(("templates/*.html"))),
	}

	app.Renderer = t

	app.GET("/", home)

	app.Logger.Fatal(app.Start(":5000"))
}
