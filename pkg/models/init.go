package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func InitDB(dsn string) {
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Ping to ensure connection is valid
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	// Set connection pool settings (optional)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&Post{}, &User{}, &Comment{}, &Like{})

	log.Println("Database connection and migration successful.")

}
