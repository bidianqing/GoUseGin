package router

import (
	"github.com/bidianqing/go-use-gin/api/handler"
	"github.com/gin-gonic/gin"
)

// 注册路由
func Map(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello Gin")
	})

	var userController = handler.UserController{}
	app.GET("/users", userController.GetUserList)
}
