package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"strconv"
	"unicode/utf8"
)

type ArticlesController struct {
}

type ArticlesFormData struct {
	Title, Body string
	URL string
	Errors map[string]string
}

func validateArticleFormData(title string, body string) map[string]string {
	errors := make(map[string]string)
	// 验证标题
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度需介于 3-40"
	}

	// 验证内容
	if body == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需大于或等于 10 个字节"
	}

	return errors
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
		tmpl, err := template.New("show.tpl").Funcs(template.FuncMap{
			"Int64ToString" : types.Int64ToString,
			"RouteName2URL" : route.Name2URL,
		}).ParseFiles("resources/views/articles/show.tpl")
		logger.LogError(err)
		tmpl.Execute(w, article)
	}

}

func (ac *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {

	articles, err := article.GetAll()
	fmt.Println("文章数据：", articles)

	if err != nil {

		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	} else {

		tmpl, err := template.ParseFiles("resources/views/articles/index.tpl")
		logger.LogError(err)

		tmpl.Execute(w, articles)
	}

}

func (ac *ArticlesController) Create(w http.ResponseWriter, r *http.Request) {

	storeURL := route.Name2URL("articles.store")
	data := ArticlesFormData{
		Title:  "",
		Body:   "",
		URL:    storeURL,
		Errors: nil,
	}
	tmpl, err := template.ParseFiles("resources/views/articles/create.tpl")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, data)
}

func (ac *ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	title :=  r.FormValue("title")
	body := r.FormValue("body")

	errors := validateArticleFormData(title, body)

	if len(errors) == 0 {
		_article := article.Article{
			Title: title,
			Body:  body,
		}
		_article.Create()
		if _article.ID > 0 {
			fmt.Fprint(w, "插入成功，ID 为"+strconv.FormatInt(_article.ID, 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章失败，请联系管理员")
		}

	} else {

		storeURL := route.Name2URL("articles.store")
		data := ArticlesFormData{
			Title: title,
			Body: body,
			URL: storeURL,
			Errors: errors,
		}

		tmpl, e := template.ParseFiles("resources/views/articles/create.tpl")
		if e != nil {
			panic(e)
		}

		tmpl.Execute(w, data)
	}
}