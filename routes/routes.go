package routes

import (
	"golang-sample-crud-service/database"
	"golang-sample-crud-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book

	err := c.BindJSON(&book)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	id := database.CreateBook(book)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Book created successful", "data": id})
}
