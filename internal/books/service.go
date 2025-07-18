package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getListOfBooks(r *gin.Engine) {
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, booksList)
	})
}

func getBookByID(r *gin.Engine) {
	r.GET("/books/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, booksList)
	})
}
