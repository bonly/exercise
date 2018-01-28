package router

import (
	"home"

	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func Setup() {
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets/", "assets/")
	router.GET("/", home.IndexPage)
}

func Run() {
	manners.ListenAndServe(":8080", router)
}
