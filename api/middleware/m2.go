package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func M2(ctx *gin.Context) {
	fmt.Println("第二个自定义中间件")
}
