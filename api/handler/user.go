package handler

import (
	"fmt"
	"net/http"

	"github.com/bidianqing/go-use-gin/internal/myappname/domain/useraggregate"
	"github.com/bidianqing/go-use-gin/internal/myappname/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct{}

// 获取用户列表
func (userController UserController) GetUserList(ctx *gin.Context) {
	// 获取queryString参数
	name := ctx.Query("name")
	fmt.Println("querystring参数name=" + name)

	repos := infrastructure.NewRepositories()
	var userRepo useraggregate.UserRepo = repos.UserRepo
	users := userRepo.GetUserList()

	ctx.JSON(200, users)
}

func (userController UserController) AddUser(ctx *gin.Context) {
	var user useraggregate.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		errors := err.(validator.ValidationErrors)

		errorMessage := user.GetModelInvalidMessages()[errors[0].Field()+"."+errors[0].Tag()]
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": errorMessage, "data": nil})
		return
	}

	ctx.JSON(200, user)
}

func (userController UserController) Error(ctx *gin.Context) {
	/*
		go func() {
			// 如果协程里的代码panic了，在没有defer recover的情况下，程序会终止退出
			// 想捕获/处理 协程内panic 所造成的恐慌，recover 必须与 defer 配套使用，否则无效
			// panic 只能触发当前 协程 的 defer 调用，在 defer 调用中如果存在 recover ，那么就能够处理其所抛出的恐慌事件。但是需要注意的是在其它 Goroutine 中的 defer 是对其没有用的，并不支持跨协程（goroutine），需要分清楚。
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()

			count := 0
			fmt.Println(3 / count)
		}()
		//*/

	count := 0
	fmt.Println(3 / count)
}
