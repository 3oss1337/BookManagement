package middleware

import (
	"context"
	"net/http"
	"time"

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
	return func(c *gin.Context) {
		context, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		c.Request = c.Request.WithContext(context)
		done := make(chan bool)
		go func() {
			c.Next()
			done <- true
		}()
		select {
		case <-done:
			return
		case <-context.Done():
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request Timeout"})
			c.Abort()
		}

	}
}
