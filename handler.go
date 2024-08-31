package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("x-api-key") == "" {
			c.AbortWithStatus(http.StatusForbidden)
		}
		if c.Request.Header.Get("x-api-key") != os.Getenv("API-Key") {
			c.AbortWithStatus(http.StatusForbidden)
		}
		c.Next()
	}
}

func Receiver(c *gin.Context) {
	c.JSON(200, gin.H{})
}
