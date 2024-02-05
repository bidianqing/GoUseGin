package useraggregate

// UserRepo 用户仓储接口
type UserRepo interface {
	GetUserList() *[]User
}
