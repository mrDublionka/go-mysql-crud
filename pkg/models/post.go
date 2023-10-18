package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mrDublionka/go-mysql-crud/pkg/config"
)

type Post struct {
	gorm.Model

	ID      uint   `gorm:"primary_key;auto_increment" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Post{})
}

func (p *Post) CreatePost() *Post {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllPosts() []Post {
	var Posts []Post
	db.Find(&Posts)
	return Posts
}

func GetPostById(Id int64) (*Post, *gorm.DB) {
	var getPost Post
	db := db.Where("ID = ?", Id).First(&getPost)
	return &getPost, db
}

func DeletePost(ID int64) Post {
	var post Post
	db.Where("ID=?", ID).Delete(post)
	return post
}
