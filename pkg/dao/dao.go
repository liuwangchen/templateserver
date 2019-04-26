package dao

import "templateserver/pkg/dao/model"

type IUserDao interface {
	GetUserById(id int) (*model.User, error)
	CreateUser(user *model.User) (int, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}
