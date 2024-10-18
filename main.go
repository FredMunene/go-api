package main

import (
	"log"
)

// POST : create
// PUT : update
// DELETE : delete
// GET : read

// methods
//

func main() {
	// router := routes.NewRouter()
	port := "localhost:3000"

	server := APIServer{
		address: port,
	}

	if err := server.Start(); err != nil {
		log.Fatalf("Error starting server : %v \n", err)
	}

	// fmt.Printf("Server running at http://%s", port)

	// err := http.ListenAndServe(port, router)
	// if err != nil {
	// 	log.Fatalf("Error starting server : %v \n", err)
	// }
}

// // get - query parameter :: /blogs?title=best
