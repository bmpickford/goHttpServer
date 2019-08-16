package main

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	ID       int
	Username string `json:"username"`
}

func HandleCustomerRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		customerPost(r)
	}
	w.WriteHeader(200)
}

func customerPost(r *http.Request) {
	var customer Customer

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&customer)
	if err != nil {
		panic(err)
	}

	err = SaveCustomer(customer.Username)
	if err != nil {
		panic(err)
	}
}
