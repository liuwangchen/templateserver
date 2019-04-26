package impl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"templateserver/pkg/dao/model"
)

type UserController struct {
	BaseController
}

func (*UserController) CreateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	_, err = userDao.CreateUser(&user)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, "create success")
	}
}

func (*UserController) UpdateUser(c *gin.Context) {
	var user model.User
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	user.ID = uint(id)
	err = c.ShouldBind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	err = userDao.UpdateUser(&user)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, "update success")
	}
}

func (*UserController) DeleteUser(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = userDao.DeleteUser(id)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, "delete success")
	}
}

func (*UserController) GetUserById(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, err := userDao.GetUserById(id)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.JSON(http.StatusOK, user)
	}
}
