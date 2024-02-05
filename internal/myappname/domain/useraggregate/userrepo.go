package useraggregate

type UserRepo interface {
	GetUserList() *[]User
}
