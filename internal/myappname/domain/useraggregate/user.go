package useraggregate

import (
	"time"

	"github.com/bidianqing/go-use-gin/internal/myappname/domain/seedwork"
)

// User 用户
type User struct {
	seedwork.Entity
	Id         int64
	Name       string
	CreateTime time.Time
}