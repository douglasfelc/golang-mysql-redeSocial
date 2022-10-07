package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// LoadTemplates load the html templates in the templates variable
func LoadTemplates() {
	// Where are the template files located
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecuteTemplate renders an html page on the screen
func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
