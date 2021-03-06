{{define "sidebar"}}
<div class="col-md-3 blog-sidebar">
    <div class="p-4 mb-3 bg-white rounded shadow-sm">
        <h1>GoBlog</h1>
        <p class="mb-0">无框goLang架版</p>
    </div>

    <div class="p-4 bg-white rounded shadow-sm mb-3">
        <h5>分类</h5>
        <ol class="list-unstyled mb-0">
            {{ range $key, $category := .Categories }}
            <li><a href="{{ $category.Link }}">{{ $category.Name }}</a></li>
            {{ end }}
            <li><a href="{{ RouteName2URL "categories.create" }}">+ 新建分类</a></li>
        </ol>
    </div>

    <div class="p-4 bg-white rounded shadow-sm mb-3">
        <h5>作者</h5>
        <ol class="list-unstyled mb-0">
            <li><a href="#">halweg</a></li>
            <li><a href="#">Aufree</a></li>
            <li><a href="#">Monkey</a></li>
        </ol>
    </div>

    <div class="p-4 bg-white rounded shadow-sm mb-3">
        <h5>链接</h5>
        <ol class="list-unstyled">
            <li><a href="#">关于我们</a></li>
            {{ if .isLogined }}
            <li><a href="{{ RouteName2URL "articles.create" }}">开始写作</a></li>
            <li class="mt-3">
                <form action="{{ RouteName2URL "auth.logout" }}" method="POST" onsubmit="return confirm('您确定要退出吗？');">
                <button class="btn btn-block btn-outline-danger btn-sm" type="submit" name="button">退出</button>
                </form>
            </li>
            {{ else }}
            <li><a href="{{ RouteName2URL "auth.register" }}">注册</a></li>
            <li><a href="{{ RouteName2URL "auth.login" }}">登录</a></li>
            {{ end }}
        </ol>
    </div>
</div>
{{end}}