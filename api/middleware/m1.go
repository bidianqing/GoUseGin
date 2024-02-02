package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func M1(ctx *gin.Context) {
	fmt.Println("第一个自定义中间件")
}
