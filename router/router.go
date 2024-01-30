package router

import (
	"github.com/bidianqing/go-use-gin/controllers"
	"github.com/gin-gonic/gin"
)

// 注册路由
func Route(app *gin.Engine) {
	var userController = controllers.UserController{}
	app.GET("/users", userController.GetUserList)
}
