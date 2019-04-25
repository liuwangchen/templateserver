package impl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (*UserController) GetUserById(c *gin.Context) {
	userid := c.Param("userid")
	c.JSON(http.StatusOK, gin.H{
		"get": fmt.Sprintf("graceful + %v", userDao.GetUserById(userid)),
	})
}
