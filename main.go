package main

import (
	. "css/db"
	. "css/middlewares"
	. "css/render"
	. "css/route"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()
	// 加载session
	// 创建基于cookie的存储引擎，yuan 参数是用于加密的密钥，可以随便填写
	store := cookie.NewStore([]byte("oliver"))
	// 设置session中间件，参数mySession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎
	r.Use(sessions.Sessions("sessionID", store))
	//加载中间件
	r.Use(PermissionMD())
	r.Static("static", "./static")
	//自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"Add": Add,
		"In":  In,
	})
	r.HTMLRender = CreateMyRender(r)

	//路由初始化
	InitRouter(r)
	//数据库初始化
	InitMySQL()
	//InitData()
	r.Run()
}
