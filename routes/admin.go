package routes

import (
    "github.com/gorilla/mux"
    "goblog/app/http/controllers/admin_controllers"
    "goblog/app/http/middlewares"
)

func RegisterAdminRoutes(r *mux.Router)  {


	auc := new(admin_controllers.AuthController)
	r.HandleFunc("/admin/auth/login", auc.Login).Methods("GET").Name("admin.auth.login")
	r.HandleFunc("/admin/auth/do-login", auc.DoLogin).Methods("POST").Name("admin.auth.dologin")

    adminRoutes := r.PathPrefix("/admin").Subrouter()
    adminRoutes.Use(middlewares.AdminAuth)

    dashController := new(admin_controllers.DashboardController)
	adminRoutes.HandleFunc("/dash-board/layout", dashController.Layout).Methods("GET").Name("admin.dash-board.layout")

}
