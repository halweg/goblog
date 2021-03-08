package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers/admin_controllers"
)

func RegisterAdminRoutes(r *mux.Router)  {

	adminRoutes := r.PathPrefix("/admin").Subrouter()

	auc := new(admin_controllers.AuthController)
	adminRoutes.HandleFunc("/auth/login", auc.Login).Methods("GET").Name("admin.auth.login")


	dashController := new(admin_controllers.DashboardController)
	adminRoutes.HandleFunc("/dash-board/layout", dashController.Layout).Methods("GET").Name("admin.dash-board.layout")

}
