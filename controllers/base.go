package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

// 获取用户列表
func (baseController BaseController) Success(ctx *gin.Context) {
	fmt.Println("获取用户列表")
}
