package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model

	ID       uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UserID   uint      `json:"user_id"`
	Image    string    `json:"image"`
	Likes    []Like    `gorm:"foreignkey:PostID" json:"likes"`
	Comments []Comment `gorm:"foreignkey:PostID" json:"comments"`
	Topic    string    `json:"topic"`
	Date     uint      `json:"date"`
	Deleted  bool      `json:"deleted" default:"false"`
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

func GetPostId(id int64) (*Post, error) {
	var post Post
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func DeletePost(db *gorm.DB, ID int64) Post {
	var post Post
	db.Where("ID=?", ID).Delete(&post)
	return post
}
