package dao

type IUser interface {
	GetUserById(id string) string
}
