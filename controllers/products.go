package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"web/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Index", models.ListAllProducts())
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")

		price, errPrice := strconv.ParseFloat(r.FormValue("price"), 64)
		if errPrice != nil {
			log.Println("Price value is invalid.", errPrice)
		}

		amount, errAmount := strconv.Atoi(r.FormValue("amount"))
		if errAmount != nil {
			log.Println("Amount value is invalid.", errAmount)
		}

		models.CreateProduct(name, description, price, amount)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Id value is invalid.", err)
	}

	models.RemoveProduct(id)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
