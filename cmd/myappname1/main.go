package main

import (
	"flag"

	router "github.com/bidianqing/go-use-gin/api"
	"github.com/bidianqing/go-use-gin/api/middleware"
	config "github.com/bidianqing/go-use-gin/configs"
	"github.com/bidianqing/go-use-gin/internal/pkg/redis"
	static "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// 获取运行环境 Development  Staging or Production
	var environmentName string
	flag.StringVar(&environmentName, "env", "Development", "环境变量")

	flag.Parse()
	isProduction := environmentName == "Production"
	if isProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()

	config.Load(app, environmentName)

	redis.NewRedisClient(config.GetString("ConnectionStrings.Redis"))
	defer redis.CloseRedisClient()

	app.Use(static.Serve("/", static.LocalFile("./wwwroot", false)))
	if !isProduction {
		app.Use(gin.Logger())
	}
	app.Use(middleware.ExceptionHandler)

	router.Map(app)

	app.Run()
}
