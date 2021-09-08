package controllers

import (
	"log"
	"net/http"

	"br.com.teste/config/database"
	"br.com.teste/models/book"
	"br.com.teste/service"
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

func FindAllBook(c *gin.Context) {
	books, err := service.FindAllBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error when find books.")
	} else {
		c.JSON(http.StatusOK, books)
	}
}
