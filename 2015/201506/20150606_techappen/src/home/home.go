package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"main.html",
		gin.H{
			"title": "千道科技",
		},
	)
}
