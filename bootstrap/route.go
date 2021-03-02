package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/routes"
)

func SetupRoute() *mux.Router {
	r := mux.NewRouter()
	routes.RegisterWebRoutes(r)
	return r
}