package controllers

import (
    "fmt"
    "goblog/pkg/flash"
    "goblog/pkg/logger"
    "gorm.io/gorm"
    "net/http"
)

type BaseController struct {
}

func (bc *BaseController) ResponseForSQLError(w http.ResponseWriter, err error) {
    if err == gorm.ErrRecordNotFound {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprint(w, "404 文章未找到")
    } else {
        logger.LogError(err)
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "服务器内部错误!")
    }
}

func (bc BaseController) ResponseForUnauthorized(w http.ResponseWriter, r *http.Request) {
    flash.Warning("未授权操作！")
    http.Redirect(w, r, "/", http.StatusFound)
}
