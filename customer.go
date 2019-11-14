package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type customer struct {
	ID       int
	Username string `json:"username"`
}

func (s *server) GetCustomer() httprouter.Handle {
	type Req struct {
		ID       int
		Username string `json:"username"`
	}
	type Res struct {
		Customer *customer
	}

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var req Req

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}
		customer, err := s.db.GetCustomer(req.Username)
		if err != nil {
			panic(err)
		}

		res := &Res{
			Customer: customer,
		}

		encoder := json.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			panic(err)
		}
	}
}

func (s *server) PostCustomer() httprouter.Handle {
	type Req struct {
		ID       int
		Username string `json:"username"`
	}
	type Res struct {
		Error string
	}

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var req Req

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		err = s.db.SaveCustomer(req.Username)
		if err != nil {
			panic(err)
		}

		res := &Res{
			Error: "",
		}
		fmt.Fprintf(w, res.Error)
	}
}
