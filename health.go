package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HealthCheck ...
func (s *server) HealthCheck() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		log.Println("Health endpoint hit")
		fmt.Fprintf(w, "healthy")
	}
}
