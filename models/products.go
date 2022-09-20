package models

import (
	"web/db"
)

type Product struct {
	Id, Amount        int
	Name, Description string
	Price             float64
}

func ListAllProducts() []Product {
	db := db.DatabaseConnect()

	selectedProductsList, err := db.Query("select * from products")
	if err != nil {
		panic(err)
	}

	products := []Product{}

	for selectedProductsList.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectedProductsList.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err)
		}

		products = append(products, Product{id, amount, name, description, price})
	}

	defer db.Close()
	return products
}

func CreateProduct(name string, description string, price float64, amount int) {
	db := db.DatabaseConnect()

	insertScript, err := db.Prepare("insert into products(name, description, price, amount) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err)
	}

	insertScript.Exec(name, description, price, amount)
	defer db.Close()
}

func RemoveProduct(id int) {
	db := db.DatabaseConnect()

	insertScript, err := db.Prepare("delete from products where id = $1")
	if err != nil {
		panic(err)
	}

	insertScript.Exec(id)
	defer db.Close()
}
