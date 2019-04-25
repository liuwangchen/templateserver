package controller

import (
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetUserById(c *gin.Context)
}
