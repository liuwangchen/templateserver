package impl

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"templateserver/pkg/dao/model"
)

var (
	db *gorm.DB
)

func init() {
	dbmysql, err := gorm.Open("mysql", "root:123456@/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db = dbmysql
	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
	}
}
