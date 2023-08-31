package core

import (
	. "css/db"
	. "css/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func GetAllTeacher(ctx *gin.Context) {
	var teachers []Teacher
	DB.Find(&teachers)
	ctx.HTML(200, "teacher.html", gin.H{
		"teachers": teachers,
	})
}

func AddTeacherHtml(ctx *gin.Context) {

	ctx.HTML(200, "addTeacher.html", gin.H{})
}

func AddTeacher(ctx *gin.Context) {
	name := ctx.PostForm("name")
	tel := ctx.PostForm("tel")
	pwd := ctx.PostForm("pwd")
	birthStr := ctx.PostForm("birth")
	loc, _ := time.LoadLocation("Asia/Shanghai")
	birth, _ := time.ParseInLocation("2006-01-02", birthStr, loc)

	var role Role
	DB.Where("name=?", "老师").Find(&role)
	user := User{
		Account: name,
		Pwd:     pwd,
		Role:    role,
	}
	NewTeacher := Teacher{
		BaseModel: BaseModel{Name: name},
		Tel:       tel,
		Birth:     &birth,
		User:      user,
	}
	fmt.Println("teacher::::::", NewTeacher)
	DB.Create(&NewTeacher)
	ctx.Redirect(301, "/teacher")
}

func EditTeacherHtml(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	var teacher Teacher
	DB.Preload("User").Where("id=?", id).Take(&teacher)
	ctx.HTML(200, "editTeacher.html", gin.H{
		"teacher": teacher,
	})
}
func EditTeacher(ctx *gin.Context) {
	fmt.Println("#################")
	id, _ := ctx.Params.Get("id")
	name := ctx.PostForm("name")
	tel := ctx.PostForm("tel")
	birthStr := ctx.PostForm("birth")
	loc, _ := time.LoadLocation("Asia/Shanghai")
	birth, _ := time.ParseInLocation("2006-01-02", birthStr, loc)
	pwd := ctx.PostForm("pwd")

	fmt.Println("pwd:::::", pwd)
	DB.Model(User{}).Where("account=?", name).Update("Pwd", pwd)
	DB.Model(Teacher{}).Where("id=?", id).Updates(map[string]interface{}{
		"Name":  name,
		"Tel":   tel,
		"Birth": birth,
	})
	ctx.Redirect(301, "/teacher")
}
func DeleteTeacher(ctx *gin.Context) {
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	id, ok := ctx.Params.Get("id")
	fmt.Println("ok:::::", ok)
	if ok {
		fmt.Println("id::::::", id)
		DB.Select("Class").Where("id=?", id).Delete(Teacher{})
		ctx.Redirect(301, "/teacher")
	} else {
		ctx.String(200, "无效的id！")
	}

}
