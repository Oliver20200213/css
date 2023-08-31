package middlewares

import (
	. "css/db"
	. "css/model"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"regexp"
	"strconv"
	"strings"
)

func PermissionMD() gin.HandlerFunc {
	return func(context *gin.Context) {
		//1.获取请求路径
		path := context.Request.URL.Path
		//2.校验路径是否在白名单内（访问在白名单内的路径直接放行）
		validUrlSlice := []string{"/login", "/reg", "^/static/.*", "/logout"}
		for _, validUrl := range validUrlSlice {
			//使用正则来验证匹配路径
			//根据正则来制定规则re
			re := regexp.MustCompile(validUrl)
			//使用re来校验path符不符合，返回的ret是一个切片
			ret := re.FindAllStringSubmatch(path, -1)
			//如果ret不等于0说明匹配成功，进行下一个验证
			if len(ret) != 0 {
				context.Next()
				return
			}
		}
		//3.登陆校验，不在路径白名单内需要登陆才能访问
		session := sessions.Default(context)
		userID := session.Get("user_id")
		if userID == nil {
			context.Redirect(301, "/login")
			return
		}

		//4.权限校验
		var user User
		DB.Where("id=?", userID).Preload("Role.Permissions").Take(&user)
		if user.RoleID == 1 {
			//取出当前登陆的学生对象，并获取sno
			var student Student
			DB.Where("user_id = ?", userID).Preload(clause.Associations).Take(&student)
			snoStr := strconv.Itoa(student.Sno)
			fmt.Println("STUDENT:::::", student)

			//循环进行权限的匹配
			for _, permission := range user.Role.Permissions {
				permissionReg := strings.Replace(permission.Url, "\\d+", snoStr, -1)
				fmt.Println("PATH:::::", path)
				fmt.Println("PERMISSIONURL::::", permissionReg)
				re := regexp.MustCompile("^" + permissionReg + "$")
				ret := re.FindAllStringSubmatch(path, -1)
				fmt.Println("ret::::::", ret)
				if len(ret) != 0 {
					//匹配成功 放行
					// 将登陆对象存储到context上下文中，方便随时取用改对象
					//主要是用于显示不同的页面
					context.Keys["loginUser"] = student
					context.Next()
					return
				}
			}

		} else {
			fmt.Println("admin role.....")
			//登陆人为管理员
			var admin Admin
			DB.Where("user_id=?", userID).Preload(clause.Associations).Take(&admin)
			fmt.Println("admin:::", admin)

			for _, permission := range user.Role.Permissions {
				re := regexp.MustCompile("^" + permission.Url + "$")
				results := re.FindAllStringSubmatch(path, -1)
				fmt.Println("results::::", results)
				if len(results) != 0 {
					context.Keys["loginUser"] = admin
					context.Next()
					return
				}
			}
		}

		context.String(403, "Forbidden!!!!")
		context.Abort()

	}
}
