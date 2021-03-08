package admin_controllers

import (
	"goblog/pkg/view/admin_view"
	"net/http"
)

type DashboardController struct {
}

func (*DashboardController) Layout(w http.ResponseWriter, r *http.Request) {

	admin_view.Render(w, admin_view.D{}, "dash_board.layout")

}
