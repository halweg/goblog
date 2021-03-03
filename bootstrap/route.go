package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/pkg/route"
	"goblog/routes"
)

func SetupRoute() *mux.Router {
	r := mux.NewRouter()
	routes.RegisterWebRoutes(r)
	route.SetRoute(r)

	return r
}