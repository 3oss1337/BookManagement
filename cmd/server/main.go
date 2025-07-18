package server

import (
	"fmt"

	"github.com/3oss1337/BookManagementAPI/internal/books"
	"github.com/3oss1337/BookManagementAPI/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	router.Use(middleware.RequestLogger())
	router.Use(middleware.Timeout())
	books.RegisterRoutes(router)
	fmt.Println("Starting API Server on Port 8000")
	router.Run(":8000")
}
