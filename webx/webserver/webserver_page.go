package webserver

import (
	"errors"
	"github.com/labstack/echo"
	"html/template"
	"io"
)

// Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found " + name)
		return err
	}
	//return tmpl.Execute(w, data)
	return tmpl.Execute(w, data)
}

type SiteRegistry struct {
	BaseWebPage BaseWebPage
	WebPages    []WebPage
}

type WebPageAPI struct {
	URL                  string
	Handler              echo.HandlerFunc
	Method               string
	ServerAPIMiddleWares []string
	MiddleWares          []echo.MiddlewareFunc
	SkipDefaultServerAPIMiddleWares bool
}

type WebPage struct {
	RequireBase bool
	//Template Name
	Name string
	//Template Files
	TemplateFiles         [] string
	URL                   string
	URLs                  [] string
	Method                string
	PageHandler           echo.HandlerFunc
	MiddleWares           []echo.MiddlewareFunc
	ServerPageMiddleWares []string
	SkipDefaultServerAPIMiddleWares bool
	PageAPIs              []WebPageAPI
}

type BaseWebPage struct {
	//Template Name
	Name string
	//Template files
	TemplateFiles []string
	PageAPIs      []WebPageAPI
}
