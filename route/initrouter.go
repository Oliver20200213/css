package route

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	InitUserRouter(r)

	InitClassRouter(r)
	InitCourseRouter(r)
	//InitTeacherRouter(r)
	InitStudentRouter(r)
}
