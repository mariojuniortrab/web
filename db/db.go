package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DatabaseConnect() *sql.DB {
	conn := "user=postgres dbname=teste2 password=q1w2e3r4 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}
	return db
}
