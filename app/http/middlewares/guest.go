package middlewares

import (
    "goblog/pkg/auth"
    "goblog/pkg/flash"
    "net/http"
)

// Auth 登录用户才可访问
func Guest(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        if auth.Check() {
            flash.Warning("您已登录，无需访问")
            http.Redirect(w, r, "/", http.StatusFound)
            return
        }

        next(w, r)
    }
}
