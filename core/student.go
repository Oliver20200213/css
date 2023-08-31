package core

import (
	. "css/db"
	. "css/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func GetAllStudent(ctx *gin.Context) {
	var students []Student
	DB.Preload("Class").Preload("Courses").Find(&students)
	ctx.HTML(200, "student.html", gin.H{
		"students":  students,
		"loginUser": ctx.Keys["loginUser"],
	})
}

func AddStudentHtml(ctx *gin.Context) {
	var classes []Class
	DB.Find(&classes)
	ctx.HTML(301, "addStudent.html", gin.H{
		"classes":   classes,
		"loginUser": ctx.Keys["loginUser"],
	})
}
func AddStudent(ctx *gin.Context) {
	sno, _ := strconv.Atoi(ctx.PostForm("sno"))
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd")
	gender, _ := strconv.Atoi(ctx.PostForm("gender"))
	birthStr := ctx.PostForm("birth")
	loc, _ := time.LoadLocation("Asia/Shanghai")
	birth, _ := time.ParseInLocation("2006-01-02", birthStr, loc)
	classID, _ := strconv.Atoi(ctx.PostForm("class_id"))
	var role Role
	DB.Where("name=?", "学生").Take(&role)

	user := User{
		Account: name,
		Pwd:     pwd,
		Role:    role,
	}

	newStudent := Student{
		BaseModel: BaseModel{Name: name},
		Sno:       sno,
		Gender:    uint8(gender),
		Birth:     &birth,
		ClassID:   classID,
		User:      user,
	}
	DB.Create(&newStudent)
	ctx.Redirect(301, "/student")
}

func GetStudent(ctx *gin.Context) {
	sno := ctx.Param("sno")
	var student Student
	DB.Preload("Class").Preload("Courses.Teacher").Where("sno=?", sno).Take(&student)
	//fmt.Println("student-courses::::", student.Courses)
	//fmt.Println("student", student)
	ctx.HTML(200, "studentDetail.html", gin.H{
		"student":   student,
		"loginUser": ctx.Keys["loginUser"],
	})

}

func SelectCourseHtml(ctx *gin.Context) {
	var courses []Course
	DB.Preload("Teacher").Find(&courses)
	//fmt.Println("courses:::::", courses)

	sno := ctx.Param("sno")
	var student Student
	DB.Where("sno=?", sno).Take(&student)
	//fmt.Println("student::::", student)

	var chooseCourses []int
	//查询出学生所选课程的id 并存储到chooseCourse切片中
	DB.Select("id").Model(&student).Association("Courses").Find(&chooseCourses)

	ctx.HTML(200, "selectCourse.html", gin.H{
		"courses":      courses,
		"chooseCourse": chooseCourses,
		"student":      student,
		"loginUser":    ctx.Keys["loginUser"],
	})
}

func SelectCourse(ctx *gin.Context) {
	//获取选课学生
	sno := ctx.Param("sno")
	var student Student
	DB.Where("sno=?", sno).Take(&student)
	//获取选中课程对象
	courses := ctx.PostFormArray("courses")
	var chooseCourses []Course
	DB.Where("id in ?", courses).Find(&chooseCourses)

	DB.Model(&student).Association("Courses").Append(chooseCourses)

	ctx.Redirect(301, "/student/"+sno+"/selectCourse")

}

func CancelCourse(ctx *gin.Context) {
	// 获取学生和取消课程的id
	sno := ctx.Param("sno")
	var student Student
	DB.Where("sno = ?", sno).Take(&student)
	courseID, _ := strconv.Atoi(ctx.Param("courseID"))

	//获取要取消的课程
	var cancelCourse Course
	DB.Where("id=?", courseID).Take(&cancelCourse)
	fmt.Println("cancelCourse:::::", cancelCourse)

	//取消课程
	DB.Model(&student).Association("Courses").Delete(cancelCourse)
	ctx.Redirect(301, "/student/"+sno+"/selectCourse")

}

func EditStudentHtml(ctx *gin.Context) {
	sno := ctx.Param("sno")
	var student Student
	DB.Preload("User").Where("sno=?", sno).Take(&student)

	var classes []Class
	DB.Find(&classes)

	ctx.HTML(200, "editStudent.html", gin.H{
		"classes":   classes,
		"student":   student,
		"loginUser": ctx.Keys["loginUser"],
	})
}

func EditStudent(ctx *gin.Context) {
	UpdateSno := ctx.Param("sno")
	name := ctx.PostForm("name")
	sno, _ := strconv.Atoi(ctx.PostForm("sno"))
	pwd := ctx.PostForm("pwd")
	gender, _ := strconv.Atoi(ctx.PostForm("gender"))
	birthStr := ctx.PostForm("birth")
	loc, _ := time.LoadLocation("Asia/Shanghai")
	birth, _ := time.ParseInLocation("2006-01-02", birthStr, loc)
	classID, _ := strconv.Atoi(ctx.PostForm("class_id"))
	var updateStudent Student
	DB.Where("sno=?", UpdateSno).Take(&updateStudent)
	DB.Model(updateStudent).Updates(map[string]interface{}{
		"Name":    name,
		"Sno":     sno,
		"Gender":  uint8(gender),
		"Birth":   birth,
		"ClassID": classID,
	})
	DB.Model(User{}).Where("id=?", updateStudent.UserID).Update("Pwd", pwd)

	ctx.Redirect(301, "/student")

}

func DeleteStudent(ctx *gin.Context) {
	sno, ok := ctx.Params.Get("sno")
	if ok {
		DB.Where("sno=?", sno).Delete(Student{})
		ctx.Redirect(301, "/student")
	} else {
		ctx.String(200, "无效的学号！")
	}

}
