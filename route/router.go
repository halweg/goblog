package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

var Router *mux.Router

func Initialize() {
	Router = mux.NewRouter()
}


func Name2URL(routeName string, pairs ...string) string {
	URL, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		//checkError(err)
		return ""
	}
	return URL.String()
}


func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
