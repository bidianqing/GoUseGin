package infrastructure

import (
	"fmt"

	"github.com/bidianqing/go-use-gin/internal/myappname/domain/useraggregate"
)

type UserRepo struct {
}

func (userRepo UserRepo) GetUserList() *[]useraggregate.User {
	fmt.Println("GetUserList")
	return new([]useraggregate.User)
}
