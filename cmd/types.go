package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Template {
	return &Template{templates: template.Must(template.ParseGlob("ui/*.html"))}
}

type Item struct {
	Id             int
	ServingSize    int
	Price          float64
	HistorialPrice float64
}

type Order struct {
	Id     int
	Items  []Item
	Method string
}
