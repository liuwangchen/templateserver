package controller

import (
	"github.com/gin-gonic/gin"
)

func InitAdminRoute(ar *gin.Engine) {
	ar.GET("/ping", Ping)
}

func InitBusRoute(br *gin.Engine) {
	br.GET("/b", B)
}
