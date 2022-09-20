package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name, Description string
	Price             float64
	Amount            float64
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "T-shirt", Description: "Blue, very beautiful", Price: 39, Amount: 10},
		{"Shoes", "Confortable", 89, 3},
		{"Fone", "working", 59, 2},
		{"New Product", "Very Good", 1, 20},
	}
	templates.ExecuteTemplate(w, "Index", products)
}
