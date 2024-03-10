package middleware

import (
	"runtime/debug"

	"github.com/bidianqing/go-use-gin/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

// 异常处理中间件 TODO 记录更详细的日志信息，包括堆栈信息
func ExceptionHandler(ctx *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			log.Logger.Errorf("%s  %s", err, string(debug.Stack()))
			ctx.JSON(500, gin.H{"success": false, "message": "程序开小差了", "data": nil})
			ctx.Abort()
		}
	}()

	ctx.Next()
}
