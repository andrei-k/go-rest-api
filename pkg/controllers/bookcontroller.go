package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/andrei-k/go-rest-api/pkg/database"
	"github.com/andrei-k/go-rest-api/pkg/models"
	"github.com/gorilla/mux"
)

func CountBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	// database.Instance.Find(&books)
	count := len(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(count)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// var books []models.Book
	// database.Instance.Find(&books)
	books := database.Books // Temp using fake data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId := mux.Vars(r)["id"]

	if !bookExists(bookId) {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Book not found")
		return
	}

	var book models.Book
	// database.Instance.First(&book, bookId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// CreateBook()
// UpdateBook()
// DeleteBook()

func bookExists(bookId string) bool {
	var book models.Book
	// database.Instance.First(&book, bookId)
	if book.ID == string(0) {
		return false
	}
	return true
}
