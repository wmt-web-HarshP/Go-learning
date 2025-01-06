package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	dsn := "root:root@tcp(127.0.0.1:3306)/simplerestapi?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to connect to the database")
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
