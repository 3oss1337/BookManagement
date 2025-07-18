package internal

import (
	"net/http"
	"strconv"

	"github.com/3oss1337/BookManagement/internal/books"
	"github.com/3oss1337/BookManagement/internal/routers"
	"github.com/gin-gonic/gin"
)

func getListOfBooks(r *gin.Engine) {
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, booksList)
	})
}

func getBookByID(r *gin.Engine) {
	r.GET("/books/:id", func(c *gin.Context) {
		var idParam string = c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid book ID"})
			return
		}
		currBook := booksList[id]

		c.JSON(http.StatusOK, currBook)
	})
}
func CreateBook(r *gin.Engine) {
	r.POST("/books", func(c *gin.Context) {
		var newBook Book
		err := c.ShouldBindJSON(&newBook)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "You didn't correctly enter a book's details"})
			return
		}
		if newBook.Title == "" || newBook.Author == "" || newBook.Year <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Title and Author must not be empty, and Year must be greater than 0"})
		}
		newBook.ID = books.NextID
		books.booksList[newBook.ID] = newBook
		books.NextID++
		c.JSON(http.StatusCreated, newBook)

	})
}
func UpdateBook(r *gin.Engine) {
	r.PUT("/books/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Incorrect book ID"})
			return
		}
		var newBookDetails Book
		err1 := c.ShouldBindJSON(&newBookDetails)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid JSON format details"})
			return
		}
		if newBookDetails.Title == "" || newBookDetails.Author == "" || newBookDetails.Year <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Title and Author must not be empty, and Year must be greater than 0"})
		}
		var currBook Book = &booksList[id]
		currBook.Author = newBookDetails.Author
		currBook.Title = newBookDetails.Title
		currBook.Year = newBookDetails.Year

	})
}

func DeleteBook(r *gin.Engine) {

}
