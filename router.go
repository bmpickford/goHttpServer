package main

import "github.com/julienschmidt/httprouter"

func (s *server) setupRoutes() {
	s.router = httprouter.New()

	s.router.GET("/health", s.HealthCheck())
	s.router.GET("/customer", s.GetCustomer())
	s.router.POST("/customer", s.PostCustomer())
}
