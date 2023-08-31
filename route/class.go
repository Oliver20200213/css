package route

import (
	. "css/core"
	"github.com/gin-gonic/gin"
)

func InitClassRouter(r *gin.Engine) {
	classRouter := r.Group("/class")
	classRouter.GET("/", GetAllClass)
	classRouter.GET("/add", AddClassHtml)
	classRouter.POST("/add", AddClass)
	classRouter.GET("/edit/:id", EditClassHtml)
	classRouter.POST("/edit/:id", EditClass)
	classRouter.GET("/delete/:id", DeleteClass)
}
