package route

import (
	. "css/core"
	"github.com/gin-gonic/gin"
)

func InitCourseRouter(r *gin.Engine) {
	courseRouter := r.Group("/course")
	courseRouter.GET("/", GetAllCourse)
	courseRouter.GET("/add", AddCourseHtml)
	courseRouter.POST("/add", AddCourse)
	courseRouter.GET("/edit/:id", EditCourseHtml)
	courseRouter.POST("/edit/:id", EditCourse)
	courseRouter.GET("/delete/:id", DeleteCourse)
}
