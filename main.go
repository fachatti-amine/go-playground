package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"

	"go-todo/routes"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	bootstrapData()

	routes.RegisterRoutes(router)
	router.Run(":8080")
}

func bootstrapData() {
	fileName := "todo-data.csv"

	// Check if the file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// File doesn't exist, create a new one
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		fmt.Println("File", fileName, "created successfully!")
	} else if err != nil {
		// Handle other potential errors
		fmt.Println("Error checking file:", err)
		return
	} else {
		// File already exists
		fmt.Println("File", fileName, "already exists.")
	}

}
