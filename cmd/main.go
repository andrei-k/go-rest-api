package main

import (
	"log"
	"net/http"
	"time"

	"github.com/andrei-k/go-rest-api/pkg/controllers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register routes
	router.HandleFunc("/count", controllers.CountBooks).Methods("GET")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	// router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	// router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

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
