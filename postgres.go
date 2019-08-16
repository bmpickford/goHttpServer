package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDB() *sql.DB {
	var err error

	if db == nil {
		connStr := "user=postgres dbname=go_project"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
	}

	return db
}
