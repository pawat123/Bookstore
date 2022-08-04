package config

import (
	"Bookstore/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectSQL() {
	dsn := "root:12345@tcp(127.0.0.1:3306)/boss?parseTime=true"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(
		&model.Book{},
		&model.Member{},
		&model.Borrow{},
	)
}
