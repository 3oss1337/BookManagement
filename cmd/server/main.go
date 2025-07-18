package server

import (
	"fmt"

	"github.com/3oss1337/BookManagement/internal/middleware"
	"github.com/3oss1337/BookManagement/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	router.Use(middleware.RequestLogger())
	router.Use(middleware.Timeout())
	routers.Routers(router)
	fmt.Println("Starting API Server on Port 8000")
	router.Run(":8000")
}
