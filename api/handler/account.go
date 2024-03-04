package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccountController struct{}

// 登录
func (accountController AccountController) Login(ctx *gin.Context) {
	key := "hellojwthellocsharpbidianqingbidianqing"

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "bidianqing",
		"aud":  "golang",
		"sub":  "auth",
		"exp":  jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
		"nbf":  jwt.NewNumericDate(time.Now()),
		"iat":  jwt.NewNumericDate(time.Now()),
		"jti":  uuid.New(),
		"name": "1",
	})
	token, err := t.SignedString([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(token)

	ctx.JSON(200, gin.H{"success": true, "message": "ok", "data": token})
}
