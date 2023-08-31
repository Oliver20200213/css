package route

import (
	. "css/core"
	"github.com/gin-gonic/gin"
)

func InitStudentRouter(r *gin.Engine) {
	studentRouter := r.Group("/student")
	studentRouter.GET("/", GetAllStudent)
	studentRouter.GET("/:sno", GetStudent)
	studentRouter.GET("/add", AddStudentHtml)
	studentRouter.POST("/add", AddStudent)
	studentRouter.GET("/:sno/selectCourse", SelectCourseHtml)
	studentRouter.POST("/:sno/selectCourse", SelectCourse)
	studentRouter.GET("/:sno/cancelCourse/:courseID", CancelCourse)

	studentRouter.GET("/edit/:sno", EditStudentHtml)
	studentRouter.POST("/edit/:sno", EditStudent)
	studentRouter.GET("/delete/:sno", DeleteStudent)

}
