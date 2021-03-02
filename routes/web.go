package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"net/http"
)

// 注冊
func RegisterWebRouter(r *mux.Router)  {

	pc := new(controllers.PagesController)

	r.HandleFunc("/", pc.HomeHandler).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.AboutHandler).Methods("GET").Name("about")
	// 自定义 404 页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFoundHandler)

}
