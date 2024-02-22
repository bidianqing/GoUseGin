package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 异常处理中间件
func ExceptionHandler(ctx *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err错误信息：", err)
			ctx.JSON(500, gin.H{"success": false, "message": "程序开小差了", "data": nil})
			ctx.Abort()
		}
	}()

	ctx.Next()
}
