package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"net/http"
)

var router *mux.Router
var db *sql.DB

func main() {

	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	// 中间件：强制内容类型为 HTML
	router.Use(middlewares.ForceHTMLMiddleware)

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}