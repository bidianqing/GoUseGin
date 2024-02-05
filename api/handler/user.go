package handler

import (
	"time"

	config "github.com/bidianqing/go-use-gin/configs"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

// 获取用户列表
func (userController UserController) GetUserList(ctx *gin.Context) {
	data := make(map[string]interface{})
	data["name"] = "tom"
	data["time"] = time.Now()
	data["tags"] = []int{1, 2, 3, 4}
	data["age"] = nil
	data["appName"] = config.GetString("AppName")
	data["remote"] = config.GetString("Remote")

	ctx.JSON(200, data)
}

func (userController UserController) AddUser(ctx *gin.Context) {

}
