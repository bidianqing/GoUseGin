package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// 认证中间件
func AuthenticationHandler(ctx *gin.Context) {
	isAuthenticated := false
	authorization := ctx.GetHeader("Authorization")
	if token := strings.Trim(authorization, " "); token != "" {
		token = strings.TrimPrefix(token, "Bearer ")
		isAuthenticated = true

		fmt.Println(token)
	}

	ctx.Set("IsAuthenticated", isAuthenticated)
	if isAuthenticated {
		ctx.Set("Claims", map[string]string{
			"name": "1",
		})
	}

	ctx.Next()

}
