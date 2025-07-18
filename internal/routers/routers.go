package routers

import (
	"github.com/3oss1337/BookManagement/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.GET("/books", handlers.GetListOfBooks)
	r.GET("/books/:id", handlers.GetBookByID)
	r.POST("/books", handlers.CreateBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

}
