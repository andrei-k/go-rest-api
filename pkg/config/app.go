package config

import (
	"github.com/jinzhu/gorm"
)

var (
	// db variable that will talk to the database
	// Reference: https://gorm.io/docs/connecting_to_the_database.html
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "admin:password/books?charset=uf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
