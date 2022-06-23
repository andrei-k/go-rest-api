package database

import (
	"github.com/andrei-k/go-rest-api/pkg/models"
)

// Creates a few books as test data.
// TODO: Use a real database.
var Books = models.Books{
	models.Book{
		ID:    "1",
		Title: "On Writing Well",
		Author: models.Author{
			FirstName: "William",
			LastName:  "Zinsser",
		},
	},
	models.Book{
		ID:    "2",
		Title: "Stein on Writing",
		Author: models.Author{
			FirstName: "Sol",
			LastName:  "Stein",
		},
	},
}
