package controllers

import (
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/view"
	"gorm.io/gorm"
	"net/http"
)

type BaseController struct {
}

func (bc *BaseController) ResponseForSQLError(w http.ResponseWriter, err error) {
    if err == gorm.ErrRecordNotFound {
		logger.LogError(err)
		w.WriteHeader(http.StatusNotFound)
        view.Render(w,view.D{
        	"Title" : "资源不存在",
        	"Info" : "文章不存在",
        	"Message" : "it`s is a sad message...",
		},
        "errors.404")
    } else {
        logger.LogError(err)
        w.WriteHeader(http.StatusInternalServerError)
		view.Render(w,view.D{
			"Title" : "服务器出错了",
			"Info" : "服务器内部发生错误",
			"Message" : "it`s is a sad message...",
		},
			"errors.500")
    }
}

func (bc BaseController) ResponseForUnauthorized(w http.ResponseWriter, r *http.Request) {
    flash.Warning("未授权操作！")
    http.Redirect(w, r, "/", http.StatusFound)
}
