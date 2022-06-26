package models

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

// Defines the slice of structs holding books.
type Books []Book
