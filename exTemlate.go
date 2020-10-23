package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func exT() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("index.html")),
	}
	e.Renderer = renderer
	e.GET("/hello", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello", "Echo")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
