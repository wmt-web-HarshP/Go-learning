package models

import (
	"github.com/akhil/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(bookID uint) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", bookID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(bookID uint) Book {
	var book Book
	db.Where("ID = ?", bookID).Delete(&book)
	return book
}
