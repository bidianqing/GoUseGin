package main

import (
	"fmt"
	"reflect"

	"github.com/bidianqing/go-use-gin/controllers"
	"github.com/gin-gonic/gin"
	static "github.com/soulteary/gin-static"
)

func main() {
	app := gin.Default()
	app.Use(static.Serve("/", static.LocalFile("./wwwroot", false)))

	app.Use(func(ctx *gin.Context) {
		fmt.Println("第一个自定义中间件")
	})

	app.Use(func(ctx *gin.Context) {
		fmt.Println("第二个自定义中间件")
	})

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello Gin")
	})

	var userController = controllers.UserController{}
	var userControllerType = reflect.TypeOf(userController)
	for i := 0; i < userControllerType.NumMethod(); i++ {
		method := userControllerType.Method(i)
		fmt.Println(method.Name)
	}
	app.GET("/users", userController.GetUserList)

	app.Run(":8080")
}
