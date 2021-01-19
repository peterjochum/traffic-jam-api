package handlers

import (
	"html/template"
	"net/http"
	"path"
)

// StaticFilesPath stores the relative path to static files
var StaticFilesPath = "static"

// Index contains a welcome page that leads to the API endpoints
func Index(w http.ResponseWriter, _ *http.Request) {
	templatePath := path.Join(StaticFilesPath, "index.gohtml")
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}
	data := struct {
		Title string
	}{
		"Traffic Jam API",
	}
	if err := tpl.Execute(w, data); err != nil {
		panic(err)
	}
}
