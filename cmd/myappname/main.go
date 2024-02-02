package main

import (
	"fmt"

	"github.com/bidianqing/go-use-gin/cmd/myappname/router"
	"github.com/gin-gonic/gin"
	static "github.com/soulteary/gin-static"
)

func main() {
	app := gin.Default()
	app.Use(static.Serve("/", static.LocalFile("./cmd/myappname/wwwroot", false)))

	app.Use(func(ctx *gin.Context) {
		fmt.Println("第一个自定义中间件")
	})

	app.Use(func(ctx *gin.Context) {
		fmt.Println("第二个自定义中间件")
	})

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello Gin")
	})

	router.Map(app)

	app.Run(":8080")
}
