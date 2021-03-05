package bootstrap

import (
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/model"
	"gorm.io/gorm"
	"time"
)

func SetupDB()  {

	db := model.ConnectDB()


	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	migration(db)

}
func migration(db *gorm.DB) {
	// 自动迁移
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}

