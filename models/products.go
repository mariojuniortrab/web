package models

import "web/db"

type Product struct {
	Name, Description string
	Price             float64
	Amount            int
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

		err = selectedProductsList.Scan(&id, &name, &description, &amount, &price)
		if err != nil {
			panic(err)
		}

		products = append(products, Product{name, description, price, amount})
	}

	defer db.Close()
	return products
}
