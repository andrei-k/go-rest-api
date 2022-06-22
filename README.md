# REST API with Go

A simple CRUD REST API with Mux written in Go.

## Todo
- [x] Original implementation entirely in main.go
- [ ] Add documentation
- [ ] Optimize directory structure
- [ ] Use real database

---

## Original code entirely in main.go

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

// books is a slice of Book structs
var books []Book

// GetBooks returns all books in the books slice.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook returns a book from the books slice.
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		// Checks to see if a book matches the ID passed in as a parameter.
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// Throws and error if no books match the parameter.
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No matching book found"))
}

// CreateBook adds a new book to the books slice.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa((rand.Intn(100000)))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook updates a book in the books slice.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&books)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
		}
	}
}

// DeleteBook deletes a book from the books slice.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var params = mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
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

	// Creates a new router.
	r := mux.NewRouter()
	// Defines the routes.
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/books", CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
```
