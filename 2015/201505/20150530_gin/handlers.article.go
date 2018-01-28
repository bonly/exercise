package main

import (
	// "net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	"net/http"
)

func showIndexPage(c *gin.Context) {
	// articles := getAllArticles()

	// Call the render function with the name of the template to render
	// render(c, gin.H{
	// 	"title":   "Home Page",
	// 	"payload": articles}, "index.html")

  c.HTML(
      // Set the HTTP status to 200 (OK)
      http.StatusOK,
      // Use the index.html template
      "index.html",
      // Pass the data that the page uses (in this case, 'title')
      gin.H{
          "title": "Home Page",
      },
  )		
}