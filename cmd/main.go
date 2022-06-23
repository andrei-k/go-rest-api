package main

import (
	"log"
	"net/http"
	"time"

	"github.com/andrei-k/go-rest-api/pkg/controllers"
	"github.com/gorilla/mux"
)

func main() {
	// Initializes the router
	router := mux.NewRouter().StrictSlash(true)

	// Registers routes
	router.HandleFunc("/count", controllers.CountBooks).Methods("GET")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	// router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	// router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	server := &http.Server{
		Addr: ":8080",
		// Enforces timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      router,
	}

	log.Println("Starting server on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
