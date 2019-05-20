package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"templateserver/config"
	"templateserver/pkg/web/controller/impl"
)

var (
	userController IUserController
)

func init() {
	if config.IsMock {

	} else {
		userController = &impl.UserController{}
	}
}

func InitAdminRoute(ar *gin.Engine) {
	admin := ar.Group("/admin")
	{
		admin.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		admin.GET("/haha", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": fmt.Sprintf("%v", c.Request.URL),
			})
		})
		admin.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": c.QueryArray("p"),
			})
		})
	}
}

func InitBusRoute(br *gin.Engine) {
	userr := br.Group("/user")
	{
		userr.GET("/", userController.GetUserById)
		userr.POST("/", userController.CreateUser)
		userr.PUT("/:id", userController.UpdateUser)
		userr.DELETE("/", userController.DeleteUser)
	}
}
