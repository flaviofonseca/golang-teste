package controllers

import (
	"log"

	"br.com.teste/config/database"
	"br.com.teste/models/book"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	book := book.Book{Title: "Titulo", Author: "Author"}

	err := database.DB.Create(&book).Error

	if err != nil {
		panic("Erro ao criar um book")
	}

	log.Println("Book created", book)
	// c.JSON(http.StatusOK, gin.H{"data": book})
}
