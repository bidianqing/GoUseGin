package infrastructure

import (
	"fmt"
	"time"

	"github.com/bidianqing/go-use-gin/internal/myappname/domain/useraggregate"
	"github.com/bwmarrin/snowflake"
)

// UserRepo 用户仓储实现
type UserRepo struct {
}

func (userRepo UserRepo) GetUserList() *[]useraggregate.User {
	fmt.Println("GetUserList")

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
	}

	// Generate a snowflake ID.
	id := node.Generate()

	return &[]useraggregate.User{{Id: id.Int64(), Name: "tom", CreateTime: time.Now()}}
}
