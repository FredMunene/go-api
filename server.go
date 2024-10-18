package main

import (
	"go-api/routes"
	"log"
	"net/http"
)

type APIServer struct {
	address string
}

func newAPIServer(address string) *APIServer {
	return &APIServer{
		address: address,
	}
}

func (s *APIServer) Start() error {
	// create eouter with routes attached

	router := routes.NewRouter()

	server := http.Server{
		Addr: s.address,
		Handler: router,
	}
	log.Printf("Server started on %s", s.address)
	return server.ListenAndServe()
}
