package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", healthCheckHandler)
	router.Run(":8080")
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"healthy":       true,
		"message":       "Hello, world!",
		"message-debug": "Just changed code",
	})
}