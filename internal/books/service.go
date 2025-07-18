package books

import (
	"net/http"
	"strconv"

	"github.com/3oss1337/BookManagementAPI/internal/books"
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

		c.JSON(http.StatusOK, booksList[id])
	})
}
func CreateBook(r *gin.Engine) {
	r.POST("/books", func(c *gin.Context) {
		var newBook Book
		err := c.ShouldBindJSON(&newBook)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"Error": "You didn't correctly enter a book's details"})
		}
		newBook.ID = books.NextID
		booksList[newBook.ID] = newBook
		NextID++
		c.JSON(http.StatusCreated, newBook)

	})
}
func UpdateBook(r *gin.Engine) {

}

func DeleteBook(r *gin.Engine) {

}
