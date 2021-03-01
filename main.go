package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gorilla/mux"
)

type ArticlesFormData struct {
	Title, Body string
	URL *url.URL
	Errors map[string]string
}

/*type Article struct {
	Title, Body string
	ID          int64
}*/

type Article struct {
	Title, Body string
	ID int64
}

func (A Article) Link() string {
	URL, err := router.Get("articles.show").URL("id", strconv.FormatInt(A.ID, 10))
	if err != nil {
		checkError(err)
		return ""
	}

	return URL.String()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {

	id := getRouteVariable("id", r)

	article, err := getArticleByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误!")
		}
	} else {
		tmpl, err := template.ParseFiles("resources/views/articles/show.tpl")
		checkError(err)
		tmpl.Execute(w, article)
	}

}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {

	// 1. 执行查询语句，返回一个结果集
	rows, err := db.Query("SELECT * FROM articles")
	checkError(err)

	defer rows.Close()

	var articles []Article
	//2. 循环读取结果
	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.Title, &article.Body)
		if err != nil {
			checkError(err)
		}
		articles = append(articles, article)
	}

	// 2.3 检测遍历时是否发生错误
	if err = rows.Err(); err != nil {
		checkError(err)
	}

	// 3. 加载模板
	tmpl, err := template.ParseFiles("resources/views/articles/index.tpl")
	if err!= nil {
		checkError(err)
	}

	// 4. 渲染模板，将所有文章的数据传输进去
	tmpl.Execute(w, articles)


}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	title :=  r.FormValue("title")
	body := r.FormValue("body")

	errors := validateArticleFormData(title, body)

	if len(errors) == 0 {
		fmt.Fprintf(w, "title的内容输入合法：%v", title)
		fmt.Fprintf(w, "body的内容输入合法: %v",  body)

		lastInsertId , err := saveArticleToDB(title, body)
		if lastInsertId > 0 {
			fmt.Fprintln(w, "数据库插入成功！,id为"+strconv.FormatInt(lastInsertId, 10))
		} else {
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w,  "500 服务器内部错误")
		}

	} else {

		/*storeURL, _ := router.Get("articles.store").URL()

		data := ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    storeURL,
			Errors: errors,
		}

		tmpl, err := template.ParseFiles("resources/views/articles/create.tpl")
		if err != nil {
			panic(err)
		}

		tmpl.Execute(w, data)*/
		storeURL, _ := router.Get("articles.store").URL()
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

func forceHTMLMiddleware(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-type", "text/html; charset=utf-8")
		next.ServeHTTP(writer, request)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 除首页以外，移除所有请求路径后面的斜杆
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		// 2. 将请求传递下去
		next.ServeHTTP(w, r)
	})
}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("resources/views/articles/create.tpl")

	if err != nil {
		panic(err)
	}
	storeURL, _ := router.Get("articles.store").URL()

	data := ArticlesFormData{
		Title:  "",
		Body:   "",
		URL:    storeURL,
		Errors: nil,
	}

	tmpl.Execute(w, data)
}

func articlesEditHandler(w http.ResponseWriter, r *http.Request) {
	id := getRouteVariable("id", r)

	article, err := getArticleByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "文章未找到")
		} else {
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误")
		}
	} else {
		updateURL, _ := router.Get("articles.update").URL("id", id)

		data := ArticlesFormData{
			Title:  article.Title,
			Body:   article.Body,
			URL:    updateURL,
			Errors: nil,
		}

		tmpl, err := template.ParseFiles("resources/views/articles/edit.tpl")
		checkError(err)
		tmpl.Execute(w, data)

	}
}

func getRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

func getArticleByID(id string) (Article, error) {
	article := Article{}
	query := "SELECT * FROM articles WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)
	return article, err
}

func articlesUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := getRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_, err := getArticleByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误！")
		}
	} else {
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")

		errors := validateArticleFormData(title, body)

		if len(errors) == 0 {

			query := "UPDATE articles SET title = ?, body = ? WHERE id = ?"
			rs, err := db.Exec(query,title, body, id)
			if err != nil {
				checkError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "服务器内部错误")
			}

			if rowNum, _ := rs.RowsAffected(); rowNum > 0 {
				showURL, _ := router.Get("articles.show").URL("id", id)
				http.Redirect(w, r, showURL.String(), http.StatusFound)
			} else {
				fmt.Fprint(w, "你没有做任何修改！")
			}

		} else {

			updateURL, _ := router.Get("articles.update").URL("id", id)

			data := ArticlesFormData{
				Title:  title,
				Body:   body,
				URL:    updateURL,
				Errors: errors,
			}

			tmpl, err := template.ParseFiles("resources/views/articles/edit.tpl")
			checkError(err)

			tmpl.Execute(w, data)

		}

	}
}

func validateArticleFormData(title string, body string) map[string]string {
	errors := make(map[string]string)

	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题内容必须在3~40个字之间"
	}

	if body == "" {
		errors["body"] = "文章内容不能为空！"
	} else if utf8.RuneCountInString(body) < 10  {
		errors["body"] = "文章内容不能少于10个字！"
	}

	return errors
}

var router = mux.NewRouter().StrictSlash(true)
var db *sql.DB

func initDB() {
	var err error
	config := mysql.Config{
		User:                    "go_blog",
		Passwd:                  "go_blog123",
		Net:                     "tcp",
		Addr:                    "127.0.0.1:3306",
		DBName:                  "go_blog",
		AllowNativePasswords:    true,
	}

	db, err = sql.Open("mysql", config.FormatDSN())
	checkError(err)

	db.SetMaxOpenConns(25)

	db.SetConnMaxLifetime(5 * time.Minute)

	db.SetMaxIdleConns(25)

	err = db.Ping()
	checkError(err)
}

func checkError(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
		id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		body longtext COLLATE utf8mb4_unicode_ci
	); `

	_, err := db.Exec(createArticlesSQL)
	checkError(err)
}

func saveArticleToDB(title string, body string) (int64, error) {

	// 变量初始化
	var (
		id   int64
		err  error
		rs   sql.Result
		stmt *sql.Stmt
	)


	// 1. 获取一个 prepare 声明语句
	stmt, err = db.Prepare("INSERT INTO articles (title, body) values (?,?)")
	// 例行的错误检测
	if err != nil {
		return 0, err
	}

	// 2. 在此函数运行结束后关闭此语句，防止占用 SQL 连接
	defer stmt.Close()

	// 3. 执行请求，传参进入绑定的内容
	rs, err = stmt.Exec(title, body)
	if err != nil {
		return 0, err
	}

	// 4. 插入成功的话，会返回自增 ID
	if id, err = rs.LastInsertId(); id != 0 {
		return id, nil
	}
	return 0, err
}

func main() {
	initDB()
	createTables()

	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")

	router.HandleFunc("/articles/create", articlesCreateHandler).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/{id:[0-9]+}/edit", articlesEditHandler).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesUpdateHandler).Methods("POST").Name("articles.update")

	// 自定义 404 页面
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// 中间件：强制内容类型为 HTML
	router.Use(forceHTMLMiddleware)

	// 通过命名路由获取 URL 示例
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "1")
	fmt.Println("articleURL: ", articleURL)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}