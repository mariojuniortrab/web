package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func databasConnect() *sql.DB {
	conn := "user=postgres dbname=teste2 password=q1w2e3r4 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}
	return db
}

type Product struct {
	Name, Description string
	Price             float64
	Amount            int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := databasConnect()

	selectedProductsList, err := db.Query("select * from products")
	if err != nil {
		panic(err)
	}

	p := Product{}
	products := []Product{}

	for selectedProductsList.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectedProductsList.Scan(&id, &name, &description, &amount, &price)
		if err != nil {
			panic(err)
		}

		p.Name = name
		p.Description = description
		p.Amount = amount
		p.Price = price

		products = append(products, p)

	}

	templates.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
