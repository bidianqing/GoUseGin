package router

import (
	"github.com/bidianqing/go-use-gin/cmd/myappname/handler"
	"github.com/gin-gonic/gin"
)

// 注册路由
func Map(app *gin.Engine) {
	var userController = handler.UserController{}
	app.GET("/users", userController.GetUserList)
}
