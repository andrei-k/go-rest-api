package main

import (
	"log"
	"net/http"
	"time"

	"github.com/andrei-k/go-rest-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)
	routes.RegisterBookRoutes(router)

	server := &http.Server{
		Addr: ":8080",
		// Enforce timeouts, which is good practice
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      router,
	}

	log.Println("Starting server on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
