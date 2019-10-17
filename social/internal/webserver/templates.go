package webserver

import (
	"html/template"
	"path"
)


func NewTemplates() map[string]*template.Template {
	templates := make(map[string]*template.Template)
	for _, item := range temps() {
		templates[item.name] = template.Must(template.ParseFiles(item.base, item.child))
	}
	return templates
}

type templateName struct {
	name string
	base string
	child string
}

func temps() []templateName {
	absPath := "/code/html"
	return []templateName{
		{
			child: path.Join(absPath, "loginForm.html"),
			base: path.Join(absPath,"base.html"),
			name: "login",
		},
		{
			child: path.Join(absPath, "registrationForm.html"),
			base: path.Join(absPath,"base.html"),
			name: "registration",
		},
		{
			child: path.Join(absPath, "userProfile.html"),
			base: path.Join(absPath,"base.html"),
			name: "profile",
		},
		{
			child: path.Join(absPath, "index.html"),
			base: path.Join(absPath,"base.html"),
			name: "index",
		},
	}
}


