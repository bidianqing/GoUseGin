package main

import (
	"flag"
	"fmt"
	"os"

	router "github.com/bidianqing/go-use-gin/api"
	"github.com/bidianqing/go-use-gin/api/middleware"
	config "github.com/bidianqing/go-use-gin/configs"
	"github.com/bidianqing/go-use-gin/internal/pkg/redis"
	static "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	// 获取运行环境 Development  Staging or Production
	environmentName := getEnvironmentName()

	var flagEnvironmentName string
	flag.StringVar(&flagEnvironmentName, "env", "", "环境变量")
	flag.Parse()
	if flagEnvironmentName != "" {
		environmentName = flagEnvironmentName
	}

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

	fmt.Println("Hosting Environment: ", environmentName)
	app.Run()
}

func getEnvironmentName() string {
	var environmentName string
	data, err := os.ReadFile("./launchSettings.json")
	if err == nil {
		m := make(map[string]interface{})
		jsoniter.Unmarshal(data, &m)
		environmentName = m["GIN_ENVIRONMENT"].(string)
	}

	if environmentName == "" {
		environmentName = os.Getenv("GIN_ENVIRONMENT")
	}

	if environmentName == "" {
		environmentName = "Production"
	}

	return environmentName
}
