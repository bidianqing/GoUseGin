package handler

import (
	"time"

	config "github.com/bidianqing/go-use-gin/configs"
	logger "github.com/bidianqing/go-use-gin/internal/pkg/log"
	"github.com/bidianqing/go-use-gin/internal/pkg/queue"
	"github.com/bidianqing/go-use-gin/internal/pkg/redis"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

// 获取配置
func (testController TestController) Config(ctx *gin.Context) {
	key := ctx.Query("key")
	val := config.GetString(key)

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
	logger.Info("Info日志信息")

	ctx.JSON(200, gin.H{"success": true, "message": "ok", "data": nil})
}

// Queue
func (testController TestController) Queue(ctx *gin.Context) {
	queue.Queue(func() {
		time.Sleep(time.Second * 3)
		logger.Info("执行")
	})

	ctx.JSON(200, gin.H{"success": true, "message": "ok", "data": nil})
}
