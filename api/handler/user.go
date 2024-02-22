package handler

import (
	"fmt"
	"net/http"

	"github.com/bidianqing/go-use-gin/internal/myappname/domain/useraggregate"
	"github.com/bidianqing/go-use-gin/internal/myappname/infrastructure"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

// 获取用户列表
func (userController UserController) GetUserList(ctx *gin.Context) {
	// 获取queryString参数
	name := ctx.Query("name")
	fmt.Println("querystring参数name=" + name)

	repos := infrastructure.NewRepositories()
	var userRpo useraggregate.UserRepo = repos.UserRepo
	users := userRpo.GetUserList()

	ctx.JSON(200, users)
}

func (userController UserController) AddUser(ctx *gin.Context) {
	var user useraggregate.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}

func (userController UserController) Error(ctx *gin.Context) {
	count := 0

	fmt.Println(3 / count)
}
