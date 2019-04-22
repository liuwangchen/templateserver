package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func B(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"b": "haha",
	})
}
