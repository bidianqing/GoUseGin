package handler

import (
	config "github.com/bidianqing/go-use-gin/configs"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

// 获取配置
func (testController TestController) Config(ctx *gin.Context) {
	val := config.GetString("ConnectionStrings.Mysql")

	ctx.JSON(200, gin.H{"success": true, "message": "ok", "data": val})
}
