package route

import (
	. "css/core"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	r.GET("/", Index)
	r.GET("/login", LoginHtml)
	r.POST("/login", Login)
	r.GET("/logout", Logout)
}
