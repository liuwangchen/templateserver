package controller

import (
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetUserById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
