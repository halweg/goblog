package model

import (
	"goblog/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	
	var err error
	
	config := mysql.New(mysql.Config{
		DSN: "go_blog:go_blog123@tcp(127.0.0.1:3306)/go_blog?charset=utf8&parseTime=True&loc=Local",
	})

	DB, err := gorm.Open(config, &gorm.Config{})

	logger.LogError(err)

	return DB
	
}
