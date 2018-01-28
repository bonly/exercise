package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()
}

func render(c *gin.Context, data gin.H, templateName string) {

  switch c.Request.Header.Get("Accept") {
  case "application/json":
    // Respond with JSON
    c.JSON(http.StatusOK, data["payload"])
  case "application/xml":
    // Respond with XML
    c.XML(http.StatusOK, data["payload"])
  default:
    // Respond with HTML
    c.HTML(http.StatusOK, templateName, data)
  }

}