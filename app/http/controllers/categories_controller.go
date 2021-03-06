package controllers

import (
    "fmt"
    "goblog/app/models/category"
    "goblog/app/requests"
    "goblog/pkg/view"
    "net/http"
)

// CategoriesController 文章分类控制器
type CategoriesController struct {
    BaseController
}

// Create 文章分类创建页面
func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
    view.Render(w, view.D{}, "categories.create")
}

// Store 保存文章分类
func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request) {
    // 稍后开发
    _category := category.Category{
        Name: r.PostFormValue("name"),
    }

    errors := requests.ValidateCategoryForm(_category)

    if len(errors) == 0 {
        _category.Create()
        if _category.ID > 0 {
            fmt.Fprint(w , "恭喜你创建成功")
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprint(w, "创建失败，请联系管理员！")
        }

    } else {
        view.Render(w, view.D{
            "Errors" : errors,
            "Category" : _category,
        }, "categories.create")
    }
}

func (*CategoriesController) Show(w http.ResponseWriter, r *http.Request) {
    //
}