package render

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func CreateMyRender(r *gin.Engine) multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	render.AddFromFilesFuncs("login.html", r.FuncMap, "templates/login.html")
	render.AddFromFilesFuncs("index.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/index.html", "templates/studentMenu.html")

	render.AddFromFilesFuncs("class.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/class/class.html", "templates/studentMenu.html")
	render.AddFromFilesFuncs("addClass.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/class/addClass.html", "templates/studentMenu.html")
	render.AddFromFilesFuncs("editClass.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/class/editClass.html", "templates/studentMenu.html")

	render.AddFromFilesFuncs("course.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/course/course.html", "templates/studentMenu.html")
	render.AddFromFilesFuncs("addCourse.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/course/addCourse.html", "templates/studentMenu.html")
	render.AddFromFilesFuncs("editCourse.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/course/editCourse.html", "templates/studentMenu.html")

	render.AddFromFilesFuncs("teacher.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/teacher/teacher.html", "templates/studentMenu.html")
	render.AddFromFilesFuncs("addTeacher.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/teacher/addTeacher.html", "templates/studentMenu.html")
	render.AddFromFilesFuncs("editTeacher.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/teacher/editTeacher.html", "templates/studentMenu.html")

	render.AddFromFilesFuncs("student.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/student/student.html", "templates/studentMenu.html")
	render.AddFromFilesFuncs("addStudent.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/student/addStudent.html", "templates/studentMenu.html")
	render.AddFromFilesFuncs("studentDetail.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/student/studentDetail.html", "templates/studentMenu.html")
	render.AddFromFilesFuncs("selectCourse.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/student/selectCourse.html", "templates/studentMenu.html")
	render.AddFromFilesFuncs("editStudent.html", r.FuncMap, "templates/base.html", "templates/adminMenu.html", "templates/student/editStudent.html", "templates/studentMenu.html")

	return render
}
