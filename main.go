package main

import (
	"github.com/gorilla/mux"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"net/http"
)

var router *mux.Router

func main() {

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	// 中间件：强制内容类型为 HTML
	router.Use(middlewares.ForceHTMLMiddleware)

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}