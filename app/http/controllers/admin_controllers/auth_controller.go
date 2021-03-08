package admin_controllers

import (
	"goblog/pkg/view/admin_view"
	"net/http"
)

type AuthController struct {
}


func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	admin_view.RenderSimple(w, admin_view.D{}, "auth.login")
}
