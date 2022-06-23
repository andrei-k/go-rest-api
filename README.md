# REST API with Go

A simple REST API written in Go that performs CRUD operations for book management. This program also attempts to follow a clean folder structure  and best development practices. Gorilla/mux is used to route incoming HTTP requests to the correct method handlers.

## Todo
- [x] Original implementation entirely in main.go
- [x] Optimize directory structure
- [ ] Add documentation
- [ ] Use real database

---

## Original code that was entirely in main.go
One of my goals for this program is to organize the code into a better folder structure, rather than putting everything into a single main.go file.

```go
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book is a struct that holds a book's id, title, and author.
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"director"`
}

// Author is a struct that holds an author's first and last names.
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var books []Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		// Checks to see if a book matches the ID passed in as a parameter.
		if item.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// Throws and error if no books match the parameter.
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No matching book found"))
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa((rand.Intn(100000)))
	books = append(books, book)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook updates a book in the books slice.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&books)
			book.ID = params["id"]
			books = append(books, book)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
		}
	}
}

// DeleteBook deletes a book from the books slice.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func TotalBooks(w http.ResponseWriter, r *http.Request) {
	number := len(books)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(number)
}

func main() {
	// Create a few books to start with.
	book1 := Book{
		ID:    "1",
		Title: "On Writing Well",
		Author: &Author{
			FirstName: "William",
			LastName:  "Zinsser",
		},
	}
	book2 := Book{
		ID:    "2",
		Title: "Stein on Writing",
		Author: &Author{
			FirstName: "Sol",
			LastName:  "Stein",
		},
	}
	books = append(books, book1)
	books = append(books, book2)

	// Creates a new router
	router := mux.NewRouter()
	// Register routes
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books", CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	router.HandleFunc("/total", TotalBooks).Methods("GET")

	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

```
