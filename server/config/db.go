package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "sabio-ekuator"
)

var DB *sql.DB

func init() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	// db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/sabio-ekuator?sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("Successfully connected!")

}
