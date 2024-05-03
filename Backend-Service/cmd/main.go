package main

import (
	"log"

	"github.com/ABHI2598/Backend-Service/src/server"
)

func main() {
	// Create a new server instance
	srv := server.NewServer(":8080")

	// Start the server
	log.Fatal(srv.ListenAndServe())
}
