package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// 获取用户列表
func (userController UserController) GetUserList(ctx *gin.Context) {
	data := make(map[string]interface{})
	data["name"] = "tom"
	data["time"] = time.Now()
	data["tags"] = []int{1, 2, 3, 4}
	data["age"] = nil

	ctx.JSON(200, data)
}
