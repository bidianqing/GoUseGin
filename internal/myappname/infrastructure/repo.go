package infrastructure

type Repositories struct {
	UserRepo UserRepo
}

// NewRepositories 获取所有的Repo实现
func NewRepositories() *Repositories {
	return &Repositories{
		UserRepo: UserRepo{},
	}
}
