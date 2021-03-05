package controllers

import (
    "fmt"
    "goblog/app/models/article"
    "goblog/app/policies"
    "goblog/app/requests"
    "goblog/pkg/auth"
    "goblog/pkg/flash"
    "goblog/pkg/logger"
    "goblog/pkg/route"
    "goblog/pkg/view"
    "gorm.io/gorm"
    "net/http"
)

type ArticlesController struct {
}

type ArticlesFormData struct {
	Title, Body string
	Article     article.Article
	Errors      map[string]string
}

func (ac *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误!")
		}
	} else {
        view.Render(w, view.D{
            "Article": article,
            "CanModifyArticle": policies.CanModifyArticle(article),
        }, "articles.show", "articles._article_meta")	}

}

func (ac *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	articles, err := article.GetAll()
	if err != nil {

		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误!")
	} else {
        view.Render(w, view.D{
            "Articles": articles,
        }, "articles.index", "articles._article_meta")
	}

}

func (ac *ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
}

func (ac *ArticlesController) Store(w http.ResponseWriter, r *http.Request) {

    currentUser := auth.User()
    // 1. 初始化数据
    _article := article.Article{
        Title: r.PostFormValue("title"),
        Body:  r.PostFormValue("body"),
        UserID: currentUser.ID,
    }

    // 2. 表单验证
    errors := requests.ValidateArticleForm(_article)

    // 3. 检测错误
    if len(errors) == 0 {
        // 创建文章
        _article.Create()
        if _article.ID > 0 {
            indexURL := route.Name2URL("articles.show", "id", _article.GetStringID())
            http.Redirect(w, r, indexURL, http.StatusFound)
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprint(w, "创建文章失败，请联系管理员")
        }
    } else {
        view.Render(w, view.D{
            "Article": _article,
            "Errors":  errors,
        }, "articles.create", "articles._form_field")
    }
}

func (ac *ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_article, err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
        // 检查权限
        if !policies.CanModifyArticle(_article) {
            flash.Warning("未授权操作！")
            http.Redirect(w, r, "/", http.StatusFound)
        } else {
            // 4. 读取成功，显示编辑文章表单
            view.Render(w, view.D{
                "Article": _article,
                "Errors":  view.D{},
            }, "articles.edit", "articles._form_field")
        }
	}
}

func (ac *ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误！")
		}
	} else {

	    if !policies.CanModifyArticle(_article) {
            flash.Warning("未授权操作！")
            http.Redirect(w, r, "/", http.StatusForbidden)
        } else {

            // 4.1 表单验证
            _article.Title = r.PostFormValue("title")
            _article.Body = r.PostFormValue("body")

            errors := requests.ValidateArticleForm(_article)

            if len(errors) == 0 {

                // 4.2 表单验证通过，更新数据
                rowsAffected, err := _article.Update()

                if err != nil {
                    // 数据库错误
                    w.WriteHeader(http.StatusInternalServerError)
                    fmt.Fprint(w, "500 服务器内部错误")
                    return
                }

                // √ 更新成功，跳转到文章详情页
                if rowsAffected > 0 {
                    showURL := route.Name2URL("articles.show", "id", id)
                    http.Redirect(w, r, showURL, http.StatusFound)
                } else {
                    fmt.Fprint(w, "您没有做任何更改！")
                }
            } else {

                // 4.3 表单验证不通过，显示理由
                view.Render(w, view.D{
                    "Article": _article,
                    "Errors":  errors,
                }, "articles.edit", "articles._form_field")
            }
        }
	}
}

func (ac *ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	id :=  route.GetRouteVariable("id", r)

    _article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "文章未找到")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误")
		}
	} else {

        if !policies.CanModifyArticle(_article) {
            flash.Warning("您没有权限执行此操作！")
            http.Redirect(w, r, "/", http.StatusFound)
        } else {
            // 4. 未出现错误，执行删除操作
            rowsAffected, err := _article.Delete()

            // 4.1 发生错误
            if err != nil {
                // 应该是 SQL 报错了
                w.WriteHeader(http.StatusInternalServerError)
                fmt.Fprint(w, "500 服务器内部错误")
            } else {
                // 4.2 未发生错误
                if rowsAffected > 0 {
                    // 重定向到文章列表页
                    indexURL := route.Name2URL("articles.index")
                    http.Redirect(w, r, indexURL, http.StatusFound)
                } else {
                    // Edge case
                    w.WriteHeader(http.StatusNotFound)
                    fmt.Fprint(w, "404 文章未找到")
                }
            }
        }
	}

}