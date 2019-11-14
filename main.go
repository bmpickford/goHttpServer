package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	db     *datastore
	router *httprouter.Router
}

func main() {
	log.Println("Listening on port 8080")
	s := &server{}
	s.setupRoutes()
	s.setupDB()

	log.Fatal(http.ListenAndServe(":8080", s.router))
}
