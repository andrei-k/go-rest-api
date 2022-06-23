package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/andrei-k/go-rest-api/pkg/controllers"
	"github.com/gorilla/mux"
)

func main() {
	// Initializes database (gets fake data)
	// books := database.Books

	server := &http.Server{
		Addr: ":8080",
		// Enforces timeouts for created server (good practice)
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register routes
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	// router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	// router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/total", controllers.TotalBooks).Methods("GET")

	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
