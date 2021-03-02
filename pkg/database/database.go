package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"goblog/pkg/logger"
	"time"
)

var DB *sql.DB

func Initialize() {
	initDB()
	createTables()
}

func initDB() {
	var err error
	config := mysql.Config{
		User:                    "go_blog",
		Passwd:                  "go_blog123",
		Net:                     "tcp",
		Addr:                    "127.0.0.1:3306",
		DBName:                  "go_blog",
		AllowNativePasswords:    true,
	}

	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	DB.SetMaxOpenConns(25)

	DB.SetConnMaxLifetime(5 * time.Minute)

	DB.SetMaxIdleConns(25)

	err = DB.Ping()
	logger.LogError(err)
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
		id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		body longtext COLLATE utf8mb4_unicode_ci
	); `

	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}
