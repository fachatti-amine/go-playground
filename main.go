package main

import (
	"github.com/gin-gonic/gin"

	"go-todo/routes"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
