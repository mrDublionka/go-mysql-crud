package models

import (
	"blog-app/pkg/config"
	"github.com/jinzhu/gorm"
)

type Post struct {
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Post{})
}
