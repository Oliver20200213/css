package core

import (
	. "css/db"
	. "css/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllCourse(ctx *gin.Context) {
	var courses []Course
	DB.Preload("Teacher").Find(&courses)
	ctx.HTML(200, "course.html", gin.H{
		"courses":   courses,
		"loginUser": ctx.Keys["loginUser"],
	})
}

func AddCourseHtml(ctx *gin.Context) {
	var teachers []Teacher
	DB.Find(&teachers)

	ctx.HTML(200, "addCourse.html", gin.H{
		"teachers":  teachers,
		"loginUser": ctx.Keys["loginUser"],
	})
}

func AddCourse(ctx *gin.Context) {
	name := ctx.PostForm("name")
	credit, _ := strconv.Atoi(ctx.PostForm("credit"))
	period, _ := strconv.Atoi(ctx.PostForm("period"))
	teacherID, _ := strconv.Atoi(ctx.PostForm("teacher_id"))
	newTeacher := Course{
		BaseModel: BaseModel{Name: name},
		Credit:    credit,
		Period:    period,
		TeacherID: teacherID,
	}

	DB.Create(&newTeacher)
	ctx.Redirect(301, "/course")
}

func EditCourseHtml(ctx *gin.Context) {
	var teachers []Teacher
	DB.Find(&teachers)

	id := ctx.Param("id")
	fmt.Println("id::::::::", id)
	var course Course
	DB.Where("id=?", id).Take(&course)
	fmt.Println("COURSE:::::::::::", course)

	ctx.HTML(200, "editCourse.html", gin.H{
		"teachers":  teachers,
		"course":    course,
		"loginUser": ctx.Keys["loginUser"],
	})
}

func EditCourse(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.PostForm("name")
	credit, _ := strconv.Atoi(ctx.PostForm("credit"))
	period, _ := strconv.Atoi(ctx.PostForm("period"))
	teacherID, _ := strconv.Atoi(ctx.PostForm("teacher_id"))

	DB.Model(Course{}).Where("id=?", id).Updates(map[string]interface{}{
		"Name":      name,
		"Credit":    credit,
		"Period":    period,
		"TeacherID": teacherID,
	})

	ctx.Redirect(301, "/course")
}

func DeleteCourse(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if ok {
		delID, _ := strconv.Atoi(id)
		DB.Where("id=?", delID).Delete(Course{})
		ctx.Redirect(301, "/course")
	} else {
		ctx.String(200, "无效的ID！")
	}

}
