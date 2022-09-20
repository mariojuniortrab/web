package controllers

import (
	"net/http"
	"text/template"
	"web/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Index", models.ListAllProducts())
}
