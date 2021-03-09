package middlewares

import (
    "goblog/pkg/auth"
    "goblog/pkg/auth_admin"
    "goblog/pkg/flash"
    "net/http"
)

// Auth 登录用户才可访问
func Auth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        if !auth.Check() {
            flash.Warning("登录用户才能访问此页面")
            http.Redirect(w, r, "/", http.StatusFound)
            return
        }

        next(w, r)
    }
}


func AdminAuth(next http.Handler) http.Handler  {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !auth_admin.Check() {
            flash.Warning("请先登录！")
            http.Redirect(w, r, "/admin/auth/login", http.StatusFound)
            return
        }
        next.ServeHTTP(w, r)
    })
}
