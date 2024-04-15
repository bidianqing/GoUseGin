package router

import (
	"github.com/bidianqing/go-use-gin/api/handler"
	"github.com/bidianqing/go-use-gin/api/middleware"
	"github.com/gin-gonic/gin"
)

// 注册路由
func Map(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello Gin")
	})

	// https://github.com/gin-gonic/gin/issues/3047
	apiGroup := app.Group("/").Use(middleware.AuthenticationHandler)
	loginGroup := app.Group("/login")

	var userController = handler.UserController{}
	apiGroup.GET("/users", userController.GetUserList)
	apiGroup.POST("/users", userController.AddUser)
	apiGroup.GET("/error", userController.Error)

	var accountController = handler.AccountController{}
	loginGroup.POST("", accountController.Login)

	var testController = handler.TestController{}
	apiGroup.GET("/config", testController.Config)
	apiGroup.GET("/redis", testController.Redis)
	apiGroup.GET("/log", testController.Log)
	apiGroup.GET("/queue", testController.Queue)
}
