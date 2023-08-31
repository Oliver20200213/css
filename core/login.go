package core

import (
	. "css/db"
	. "css/model"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

func LoginHtml(ctx *gin.Context) {
	ctx.HTML(200, "login.html", nil)
}

func Login(ctx *gin.Context) {
	account := ctx.PostForm("user")
	pwd := ctx.PostForm("pwd")
	var user User
	DB.Where("account=? and pwd=?", account, pwd).Take(&user)
	fmt.Println("user::::::", user)
	fmt.Println("user::::::", user.ID)

	if user.ID != 0 {
		//登陆成功
		userID := strconv.Itoa(user.ID)
		session := sessions.Default(ctx)
		session.Set("user_id", userID)
		session.Save()

		ctx.Redirect(301, "/")
	} else {
		//登陆失败
		ctx.HTML(200, "login.html", gin.H{
			"err": "用户名或密码错误！",
		})

	}
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("user_id")
	session.Save()
	ctx.Redirect(301, "/login")
}
