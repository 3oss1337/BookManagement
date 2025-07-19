package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("%s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

func Timeout() gin.HandlerFunc {
	// 	return func (c *gin.Context){
	// 		var timeout = context.WithTimeout(c.Request.Context(),2*time.Second)
	// 	}
}
