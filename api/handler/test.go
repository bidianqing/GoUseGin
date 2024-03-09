package handler

import (
	"time"

	config "github.com/bidianqing/go-use-gin/configs"
	"github.com/bidianqing/go-use-gin/internal/pkg/log"
	"github.com/bidianqing/go-use-gin/internal/pkg/redis"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

// 获取配置
func (testController TestController) Config(ctx *gin.Context) {
	val := config.GetString("ConnectionStrings.Mysql")

	ctx.JSON(200, gin.H{"success": true, "message": "ok", "data": val})
}

// Redis
func (testController TestController) Redis(ctx *gin.Context) {
	redis.Set("name", "tom", time.Second*30)
	name := redis.Get("name")

	ctx.JSON(200, gin.H{"success": true, "message": "ok", "data": name})
}

// Log
func (testController TestController) Log(ctx *gin.Context) {
	log.Logger.Debug("Info日志信息")

	ctx.JSON(200, gin.H{"success": true, "message": "ok", "data": nil})
}
