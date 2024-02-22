package useraggregate

import (
	"time"

	"github.com/bidianqing/go-use-gin/internal/myappname/domain/seedwork"
)

// User 用户
type User struct {
	seedwork.Entity
	Id         int64     `json:"id" binding:"gt=0,max=100"`
	Name       string    `json:"name" binding:"required"`
	CreateTime time.Time `json:"createTime"`
}
