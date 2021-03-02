package bootstrap

import (
	"goblog/pkg/model"
	"time"
)

func SetupDB()  {

	db := model.ConnectDB()


	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

}
