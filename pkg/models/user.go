package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model // Embed the default fields from gorm.Model

	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserPwd      string `json:"user_pwd" gorm:"unique"`
	PostsCreated []Post `json:"posts_created"`
	Photo        string `json:"photo"`
	Popular      bool   `json:"popular" default:"false"`
}

func (u *User) CreateUser() *User {
	db.Create(&u)
	return u
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := db.Where("user_email = ?", email).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil // User not found, return nil and no error
		}
		return nil, err // Other database error occurred
	}
	return &user, nil
}

func GetUserByID(id int) (*User, error) {
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil // User not found, return nil and no error
		}
		return nil, err // Other database error occurred
	}
	return &user, nil
}

func GetAllUsers() ([]User, error) {
	var Users []User
	if err := db.Find(&Users).Error; err != nil {
		return nil, err
	}
	return Users, nil
}
