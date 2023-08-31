package core

import (
	. "css/db"
	. "css/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllClass(ctx *gin.Context) {
	var classes []Class
	DB.Preload("Tutor").Find(&classes)
	ctx.HTML(200, "class.html", gin.H{
		"classes":   classes,
		"loginUser": ctx.Keys["loginUser"],
	})
}

func AddClassHtml(ctx *gin.Context) {
	var teachers []Teacher
	DB.Find(&teachers)
	ctx.HTML(200, "addClass.html", gin.H{
		"teachers":  teachers,
		"loginUser": ctx.Keys["loginUser"],
	})
}

func AddClass(ctx *gin.Context) {
	name := ctx.PostForm("name")
	num, _ := strconv.Atoi(ctx.PostForm("num"))
	tutorID, _ := strconv.Atoi(ctx.PostForm("teacher_id"))
	NewClass := Class{
		BaseModel: BaseModel{Name: name},
		Num:       num,
		TutorID:   tutorID,
	}
	DB.Create(&NewClass)
	ctx.Redirect(301, "/class")
}

func EditClassHtml(ctx *gin.Context) {
	var teachers []Teacher
	DB.Find(&teachers)
	id := ctx.Param("id")
	//id, _ := ctx.Params.Get("id")
	fmt.Println("classID::::::::::", id)
	var updateClass Class
	DB.Where("id=?", id).Take(&updateClass)
	ctx.HTML(200, "editClass.html", gin.H{
		"updateClass": updateClass,
		"teachers":    teachers,
		"loginUser":   ctx.Keys["loginUser"],
	})
}

func EditClass(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.PostForm("name")
	num, _ := strconv.Atoi(ctx.PostForm("num"))
	tutorID, _ := strconv.Atoi(ctx.PostForm("tutor_id"))
	DB.Model(Class{}).Where("id=?", id).Updates(
		map[string]interface{}{
			"Name":    name,
			"Num":     num,
			"TutorID": tutorID,
		})
	ctx.Redirect(http.StatusMovedPermanently, "/class")

}

func DeleteClass(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if ok {
		delID, _ := strconv.Atoi(id)
		DB.Where("id=?", delID).Delete(Class{})
		ctx.Redirect(http.StatusMovedPermanently, "/class")
	} else {
		ctx.String(200, "无效的ID！")
	}
}
