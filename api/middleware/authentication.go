package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 认证中间件
func AuthenticationHandler(ctx *gin.Context) {
	key := "hellojwthellocsharpbidianqingbidianqing"

	isAuthenticated := false
	authorization := ctx.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(strings.Trim(authorization, " "), "Bearer ")
	if tokenString == "" {
		ctx.JSON(401, gin.H{"success": false, "message": "认证失败", "data": nil})
		ctx.Abort()
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		fmt.Println(err)
		ctx.JSON(401, gin.H{"success": false, "message": err.Error(), "data": nil})
		ctx.Abort()
		return
	}

	isAuthenticated = true
	ctx.Set("IsAuthenticated", isAuthenticated)
	ctx.Set("Claims", token.Claims.(jwt.MapClaims))

	ctx.Next()
}
