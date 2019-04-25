package impl

import (
	"templateserver/config"
	"templateserver/pkg/dao"
	"templateserver/pkg/dao/impl"
)

var (
	userDao dao.IUserDao
)

func init() {
	if config.IsMock {

	} else {
		userDao = &impl.UserDao{}
	}
}
