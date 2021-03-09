package admin_controllers

import (
    "goblog/pkg/auth_admin"
    "goblog/pkg/view/admin_view"
    "net/http"
)

type AuthController struct {
}


func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	admin_view.RenderSimple(w, admin_view.D{}, "auth.login")
}

func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request){
    username := r.PostFormValue("username")
    password := r.PostFormValue("password")

    // 2. 尝试登录
    if err := auth_admin.Attempt(username, password); err == nil {
        // 登录成功
        //flash.Success("欢迎回来！")
        http.Redirect(w, r, "/admin/dash-board/layout", http.StatusFound)
    } else {
        // 3. 失败，显示错误提示
        admin_view.RenderSimple(w, admin_view.D{
            "Error":    err.Error(),
            "UserName": username,
            "Password": password,
        }, "auth.login")
    }
}
