package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
}

func GenerateDatabase() *Database {

	hostname := os.Getenv("DB_PORT_5432_TCP_ADDR")
	port := os.Getenv("DB_PORT_5432_TCP_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	db, err := sql.Open("postgres", "postgres://"+username+":"+password+"@"+hostname+":"+port+"/helpnumber?sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic("Not ping available")
	}

	return &Database{db}
}

func (d Database) IsCountryExist(code *string) bool {
	err := d.Connection.QueryRow("SELECT code FROM countries WHERE code = $1", code).Scan(new(string))

	if err != nil {
		return false
	} else {
		return true
	}
}

func (d Database) IsLanguageExist(code *string) bool {
	err := d.Connection.QueryRow("SELECT code FROM languages WHERE code = $1", code).Scan(new(string))

	if err != nil {
		return false
	} else {
		return true
	}
}

func (d Database) IsCategoryExist(code *string) bool {
	err := d.Connection.QueryRow("SELECT name FROM phone_numbers_categories WHERE name = $1", code).Scan(new(string))

	if err != nil {
		return false
	} else {
		return true
	}
}
