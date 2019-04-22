package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"templateserver/pkg/dao"
	"templateserver/pkg/dao/impl"
)

var (
	userDao dao.IUser = &impl.UserDao{}
)

func GetUserById(c *gin.Context) {
	userid := c.Param("userid")
	c.JSON(http.StatusOK, gin.H{
		"get": userDao.GetUserById(userid),
	})
}
