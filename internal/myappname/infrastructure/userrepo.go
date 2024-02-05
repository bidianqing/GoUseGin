package infrastructure

import (
	"fmt"
	"time"

	"github.com/bidianqing/go-use-gin/internal/myappname/domain/useraggregate"
)

// UserRepo 用户仓储实现
type UserRepo struct {
}

func (userRepo UserRepo) GetUserList() *[]useraggregate.User {
	fmt.Println("GetUserList")
	return &[]useraggregate.User{{Id: 1, Name: "tom", CreateTime: time.Now()}}
}
