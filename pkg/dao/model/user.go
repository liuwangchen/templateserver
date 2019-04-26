package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `form:"name"`
	Age  int    `form:"age"`
}
