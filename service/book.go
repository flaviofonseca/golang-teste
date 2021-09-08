package service

import (
	"br.com.teste/config/database"
	"br.com.teste/models/book"
)

/**
It is not necessary
**/
func FindAllBook() (books []book.Book, err error) {
	err = database.DB.Find(&books).Error
	return books, err
}
