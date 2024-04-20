package routes

import (
	"github.com/gin-gonic/gin"

	"go-todo/controllers"
)

var indexController = new(controllers.IndexController)
var homeController = new(controllers.HomeController)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/", indexController.Index)
	router.GET("/home", homeController.Home)
}
