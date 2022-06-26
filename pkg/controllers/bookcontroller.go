package controllers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/andrei-k/go-rest-api/pkg/database"
	"github.com/andrei-k/go-rest-api/pkg/models"
	"github.com/gorilla/mux"
)

func CountBooks(w http.ResponseWriter, r *http.Request) {
	// var books []models.Book
	// database.Instance.Find(&books)
	count := len(database.Books)
	log.Println("Count: ", count)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(count)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// var books []models.Book
	// database.Instance.Find(&books)
	books := database.Books // Temporarily use fake data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId := mux.Vars(r)["id"]

	// var book models.Book
	// database.Instance.First(&book, bookId)
	books := database.Books // Temporarily use fake data
	for _, item := range books {
		// Check to see if a book matches the ID passed in as a parameter
		if item.ID == bookId {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// Throw and error if no book matches the ID
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Book not found"))
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// Generate a random ID (for now)
	book.ID = strconv.Itoa((rand.Intn(100000)))
	database.Books = append(database.Books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range database.Books {
		if item.ID == params["id"] {
			// Delete the element at index and preserve the order of the books slice.
			// This approach creates two slices from the original, books[:index] and books[i+index:]
			// and then joins them back together into a single slice.
			// The element at index is not included.
			database.Books = append(database.Books[:index], database.Books[index+1:]...)
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			database.Books = append(database.Books, book)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	// Throw and error if no book matches the ID
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Book not found"))
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId := mux.Vars(r)["id"]

	for index, item := range database.Books {
		if item.ID == bookId {
			database.Books = append(database.Books[:index], database.Books[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// Throw and error if no book matches the ID
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Book not found"))
}
