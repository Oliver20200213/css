package route

import (
	. "css/core"
	"github.com/gin-gonic/gin"
)

func InitTeacherRouter(r *gin.Engine) {
	teacherRouter := r.Group("/teacher")
	teacherRouter.GET("/", GetAllTeacher)
	teacherRouter.GET("/add", AddTeacherHtml)
	teacherRouter.POST("/add", AddTeacher)
	teacherRouter.GET("/edit/:id", EditTeacherHtml)
	teacherRouter.POST("/edit/:id", EditTeacher)
	teacherRouter.GET("/delete/:id", DeleteTeacher)
}
