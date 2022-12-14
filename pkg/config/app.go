package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db * gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "admin:shakeeb2010@tcp(127.0.0.1:3306)/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}