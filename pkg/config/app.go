package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	connectionString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", "root", "", "bookstore")
	database, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}