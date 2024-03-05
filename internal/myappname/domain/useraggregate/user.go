package useraggregate

import (
	"time"

	"github.com/bidianqing/go-use-gin/internal/myappname/domain/seedwork"
)

// User 用户
type User struct {
	seedwork.Entity
	Id         int64     `json:"id,string" binding:"gt=0"`
	Name       string    `json:"name" binding:"required"`
	CreateTime time.Time `json:"createTime"`
}

// 模型验证失败消息
func (user User) GetModelInvalidMessages() map[string]string {
	return map[string]string{
		"Id.gt":         "Id必须大于0",
		"Id.max":        "Id不能超过100",
		"Name.required": "Name不能为空",
	}
}
