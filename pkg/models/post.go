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

func (l *Like) CreateLike() *Like {
	db.NewRecord(l)
	db.Create(&l)
	return l
}

func GetAllPosts() []Post {
	var Posts []Post
	db.Find(&Posts)
	return Posts
}

func GetPostById(id int64) (*Post, error) {
	var post Post
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &post, nil
}

func (l *Like) CheckIfIsLiked(postID uint, userID uint) bool {
	var like Like
	if err := db.Where("post_id = ? AND user_id = ?", postID, userID).First(&like).Error; gorm.IsRecordNotFoundError(err) {
		return false
	}
	return true
}

func DeletePost(db *gorm.DB, ID int64) Post {
	var post Post
	db.Where("ID=?", ID).Delete(&post)
	return post
}
