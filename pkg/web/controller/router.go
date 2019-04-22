package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"templateserver/pkg/web/controller/user"
)

func InitAdminRoute(ar *gin.Engine) {
	ar.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func InitBusRoute(br *gin.Engine) {
	userr := br.Group("/user")
	{
		userr.GET("/:userid", user.GetUserById)
	}
}
