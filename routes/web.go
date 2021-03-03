package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"net/http"
)

// 注冊
func RegisterWebRoutes(r *mux.Router)  {

	pc := new(controllers.PagesController)

	r.HandleFunc("/", pc.HomeHandler).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.AboutHandler).Methods("GET").Name("about")
	// 自定义 404 页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFoundHandler)

	ac := new (controllers.ArticlesController)
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")

	r.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store")

}
