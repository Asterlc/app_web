package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	// user, dbname, password, host, ssl(true/false)
	con := "user=administrator dbname=produtos password=pass123456 host=localhost sslmode=disable"
	db, error := sql.Open("postgres", con)

	if error != nil {
		panic(error.Error())
	} else {
		return db
	}
}
