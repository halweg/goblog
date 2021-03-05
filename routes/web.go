package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"goblog/app/http/middlewares"
	"net/http"
)

// 注冊
func RegisterWebRoutes(r *mux.Router)  {


	//全局中间件
	r.Use(middlewares.StartSession)

	pc := new(controllers.PagesController)

	r.HandleFunc("/about", pc.AboutHandler).Methods("GET").Name("about")
	// 自定义 404 页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFoundHandler)

	ac := new (controllers.ArticlesController)
	r.HandleFunc("/",  ac.Index).Methods("GET").Name("home")

	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")

	r.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store")

	r.HandleFunc("/articles/{id:[0-9]+}/edit", ac.Edit).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Update).Methods("POST").Name("articles.update")

	r.HandleFunc("/articles/{id:[0-9]+}/delete", ac.Delete).Methods("POST").Name("articles.delete")

	// 中间件：强制内容类型为 HTML
	//r.Use(middlewares.ForceHTMLMiddleware)
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", auc.Register).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", auc.DoRegister).Methods("POST").Name("auth.doregister")

	r.HandleFunc("/auth/login",  middlewares.Guest(auc.Login)).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", middlewares.Guest(auc.DoLogin)).Methods("POST").Name("auth.dologin")
	r.HandleFunc("/auth/logout", middlewares.Auth(auc.Logout)).Methods("POST").Name("auth.logout")

    uc := new(controllers.UserController)
    r.HandleFunc("/users/{id:[0-9]+}", uc.Show).Methods("GET").Name("users.show")


    r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))


}
