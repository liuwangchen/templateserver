package controller

import (
	"github.com/gin-gonic/gin"
	"templateserver/config"
	"templateserver/pkg/web/controller/impl"
)

type IUserController interface {
	GetUserById(c *gin.Context)
}

var (
	userController IUserController
)

func init() {
	if config.IsMock {

	} else {
		userController = &impl.UserController{}
	}
}
