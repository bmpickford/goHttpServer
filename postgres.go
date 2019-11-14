package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func (s *server) setupDB() {
	log.Println("Setting up db....")
	var err error
	var db *sql.DB

	connStr := "user=postgres dbname=go_project"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	s.db = &datastore{
		db: db,
	}
	log.Println("Success")
}
