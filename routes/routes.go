package routes

import (
	"github.com/gin-gonic/gin"

	"go-todo/controllers"
)

var indexController = new(controllers.IndexController)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", indexController.Index)
	router.POST("/", indexController.AddItem)
}
