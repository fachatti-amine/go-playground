package controllers

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type IndexController struct{}

func (ctl IndexController) Index(context *gin.Context) {

	data := getData()

	context.HTML(
		http.StatusOK,
		"index.tmpl.html",
		gin.H{
			"title": "Insert a todo Item!",
			"data":  data,
		})
}

func (ctl IndexController) AddItem(context *gin.Context) {
	item := context.PostForm("item")
	data := getData()

	context.HTML(
		http.StatusOK,
		"index.tmpl.html",
		gin.H{
			"title": fmt.Sprintf("Last Item Inserted: %s", item),
			"data":  data,
		})
}

func getData() [][]string {
	fileName := "todo-data.csv"

	// Open the CSV file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records into a slice of slices (array of arrays)
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return data
}
