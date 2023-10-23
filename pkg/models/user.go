package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	UserID       uint   `gorm:"primary_key;auto_increment" json:"userId"`
	UserName     string `json:"userName"`
	UserEmail    string `json:"userEmail"`
	UserPwd      string `json:"userPwd"`
	Token        string `json:"token"`
	PostsCreated []Post `json:"posts_created"`
	Photo        string `json:"photo"`
	Popular      bool   `json:"popular" default:"false"`
}
