package core

import "github.com/gin-gonic/gin"

func Index(ctx *gin.Context) {

	ctx.HTML(200, "index.html", gin.H{
		"loginUser": ctx.Keys["loginUser"],
	})
}
