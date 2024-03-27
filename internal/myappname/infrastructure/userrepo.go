package infrastructure

import (
	"time"

	"github.com/bidianqing/go-use-gin/internal/myappname/domain/useraggregate"
	"github.com/bidianqing/go-use-gin/internal/pkg/idgen"
)

// UserRepo 用户仓储实现
type UserRepo struct {
}

func (userRepo UserRepo) GetUserList() *[]useraggregate.User {
	return &[]useraggregate.User{{Id: idgen.Generate(), Name: "tom", CreateTime: time.Now()}}
}
