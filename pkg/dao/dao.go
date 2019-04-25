package dao

type IUserDao interface {
	GetUserById(id string) string
}
