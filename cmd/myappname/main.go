package main

import (
	router "github.com/bidianqing/go-use-gin/api"
	"github.com/bidianqing/go-use-gin/api/middleware"
	"github.com/gin-gonic/gin"
	static "github.com/soulteary/gin-static"
)

func main() {
	app := gin.Default()
	app.Use(static.Serve("/", static.LocalFile("./cmd/myappname/wwwroot", false)))

	app.Use(middleware.M1)
	app.Use(middleware.M2)

	router.Map(app)

	app.Run(":8080")
}
