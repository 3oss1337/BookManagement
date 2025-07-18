package handlers

import (
	"net/http"
	"strconv"

	"github.com/3oss1337/BookManagement/internal/books"
	"github.com/gin-gonic/gin"
)

func GetListOfBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books.BooksList)

}

func GetBookByID(c *gin.Context) {

	var idParam string = c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid book ID"})
		return
	}
	currBook := books.BooksList[id]

	c.JSON(http.StatusOK, currBook)

}
func CreateBook(c *gin.Context) {

	var newBook books.Book
	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "You didn't correctly enter a book's details"})
		return
	}
	if newBook.Title == "" || newBook.Author == "" || newBook.Year <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Title and Author must not be empty, and Year must be greater than 0"})
	}
	newBook.ID = books.NextID
	books.BooksList[newBook.ID] = newBook
	books.NextID++
	c.JSON(http.StatusCreated, newBook)

}
func UpdateBook(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Incorrect book ID"})
		return
	}
	var newBookDetails books.Book
	err1 := c.ShouldBindJSON(&newBookDetails)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid JSON format details"})
		return
	}
	if newBookDetails.Title == "" || newBookDetails.Author == "" || newBookDetails.Year <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Title and Author must not be empty, and Year must be greater than 0"})
	}
	newBookDetails.ID = id
	books.BooksList[id] = newBookDetails
	c.JSON(http.StatusOK, newBookDetails)

}

func DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid book ID"})
		return
	}
	_, exists := books.BooksList[id]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Book not found"})
		return
	}
	delete(books.BooksList, id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted sucessfully!"})

}
