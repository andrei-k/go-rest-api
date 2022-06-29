package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// db variable that will talk to the database
	// Reference: https://gorm.io/docs/connecting_to_the_database.html
	db *gorm.DB
)

func Connect() {
	// Reference: https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dsn := "user:pass@tcp(127.0.0.1:3306)/db-books?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
