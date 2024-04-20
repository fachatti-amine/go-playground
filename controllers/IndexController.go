package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func (ctl IndexController) Index(ginContext *gin.Context) {
	ginContext.HTML(
		http.StatusOK,
		"index.tmpl.html",
		gin.H{
			"title":     "Hello, world!",
			"paragraph": "This page is rendered by go server.",
		})
}
