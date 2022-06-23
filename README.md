# REST API with Go

A simple REST API written in Go that performs CRUD operations for book management. The purpose of this program is to demonstrate a clean folder structure using best practices. Gorilla/mux is used to route incoming HTTP requests to the correct method handlers.

## Todo

- [x] Original implementation entirely in main.go
- [x] Optimize directory structure
- [ ] Add documentation
- [ ] Use real database

---

## Project structure

The project layout uses the common pattern of splitting the code into the **/cmd** and **/pkg** layout patterns.  

### `/cmd`

The `/cmd` contains the main applications for the project. If you need to have more than one application binary, the name of subdirectories should match the name of the executable application (e.g., /cmd/myapp). It's best practice not to put a lot of code in the application directory.  

### `/pkg`
The library code that can be imported and used by external projects should live in `/pkg` directory. This layout pattern allows the package to be "go gettable". which means you can use the **go get** command to fetch and install the project, its applications, and libraries (e.g., `go get github.com/andrei-k/go-rest-api/pkg`). Use caution with the code placed here because external projects will expect these libraries to work.  

### `/internal`
Your code that is not meant to be reused by external projects should live in the `/internal` directory. Go ensures that these private packages aren't not importable.  

Some notable examples of using this layout pattern include the official [Go Tools](https://github.com/golang/tools), [Kubernetes](https://github.com/kubernetes/kubernetes), and [Docker](https://github.com/docker/compose).  

Good reference for the [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

---

## Original code that was contained entirely in the main.go file

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
