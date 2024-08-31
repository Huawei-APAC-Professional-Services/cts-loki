package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("x-api-key") == "" || c.Request.Header.Get("x-api-key") != "6eb26140-66db-11ef-bc60-00155d75c0ae" {
			c.AbortWithStatus(http.StatusForbidden)
		}
		c.Next()
	}
}

func Receiver(c *gin.Context) {
	c.JSON(200, gin.H{})
}
