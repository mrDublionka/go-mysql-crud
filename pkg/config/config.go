package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/next-php-blog"

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

	fmt.Println("Connected to the MySQL database!")
}

func GetDB() *gorm.DB {
	return db
}
