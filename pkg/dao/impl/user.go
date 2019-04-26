package impl

import (
	"templateserver/pkg/dao/model"
)

type UserDao struct {
}

func (*UserDao) GetUserById(id int) (*model.User, error) {
	user := &model.User{}
	user.ID = uint(id)
	result := db.First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (*UserDao) CreateUser(user *model.User) (int, error) {
	result := db.Create(user)
	if result.Error != nil {
		return -1, result.Error
	}
	return int(user.ID), nil
}

func (*UserDao) UpdateUser(user *model.User) error {
	result := db.Model(user).Update(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (*UserDao) DeleteUser(id int) error {
	result := db.Where("id = ?", id).Delete(model.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
