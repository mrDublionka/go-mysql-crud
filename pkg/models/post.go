package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model

	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	Image     string    `json:"image"`
	Likes     []Like    `gorm:"foreignkey:PostID" json:"likes"`
	Comments  []Comment `gorm:"foreignkey:PostID" json:"comments"`
	Topic     string    `json:"topic"`
	Date      uint      `json:"date"`
	CreatedAt uint      `json:"created_at"`
	UpdatedAt uint      `json:"updated_at"`
	DeletedAt uint      `json:"deleted_at"`
	Deleted   bool      `json:"deleted" default:"false"`
}

type Like struct {
	gorm.Model

	UserID uint `json:"user_id"`
	PostID uint `json:"post_id"`
	// You can include other fields like UserID if needed
}

type Comment struct {
	gorm.Model

	PostID  uint   `json:"post_id"`
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
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
