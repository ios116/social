package webserver

import (
	"context"
	"html/template"
	"net/http"
	"path"
)

var templates map[string]*template.Template

// InitTemplate Compile view html
func (s *HttpServer) InitTemplate() map[string]*template.Template {
	templates := make(map[string]*template.Template)
	for _, item := range s.templates() {
		templates[item.name] = template.Must(template.ParseFiles(item.base, item.name))
	}
	return templates
}

type templateName struct {
	name string
	base string
}

func (s *HttpServer) templates() []templateName {

	absPath := "/code/internal/webserver/html"
	path.Join()
	return []templateName{
		{
			name: path.Join(absPath, "loginForm.html"),
			base: "base.html",
		},
	}
}

func (s *HttpServer) RenderTemplate(ctx *context.Context, w http.ResponseWriter, templateName string, date interface{}) {
	tmpl, ok := templates[templateName]
	if !ok {
		http.Error(w, "The html does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.Execute(w, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
