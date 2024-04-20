package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct{}

func (ctl HomeController) Home(ginContext *gin.Context) {
	ginContext.HTML(
		http.StatusOK,
		"home.tmpl.html",
		gin.H{
			"title":     "-> Hello, world! <-",
			"paragraph": "This page is rendered by go server.",
		})
}
