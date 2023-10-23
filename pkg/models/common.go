package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mrDublionka/go-mysql-crud/pkg/config"
)

var db *gorm.DB

func InitDB() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(
		&Post{},
		&Like{},
		&Comment{},
		&User{},
	)
}
