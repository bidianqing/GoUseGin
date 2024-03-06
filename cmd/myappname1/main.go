package main

import (
	"flag"

	router "github.com/bidianqing/go-use-gin/api"
	"github.com/bidianqing/go-use-gin/api/middleware"
	config "github.com/bidianqing/go-use-gin/configs"
	static "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// 获取运行环境 Development  Staging or Production
	var environmentName string
	flag.StringVar(&environmentName, "env", "Development", "环境变量")

	flag.Parse()

	app := gin.Default()
	config.Load(app, environmentName)

	app.Use(static.Serve("/", static.LocalFile("./cmd/myappname1/wwwroot", false)))
	app.Use(middleware.ExceptionHandler)
	app.Use(middleware.M1)

	router.Map(app)

	app.Run()
}
