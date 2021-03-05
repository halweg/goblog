{{define "title"}}
{{ .Title }}
{{end}}

{{define "main"}}
<div class="col-md-9 blog-main">

    <div class="blog-post bg-white p-5 rounded shadow mb-4">
        <h3 class="blog-post-title">{{ .Article.Title }}</h3>
        {{template "article-meta" .Article }}
        <hr>
        {{ .Article.Body }}

        {{/* 构建删除按钮  */}}
        <form class="mt-4" action="{{ RouteName2URL "articles.delete" "id" .Article.GetStringID }}" method="post">
        {{ if .CanModifyArticle }}
            <button type="submit" onclick="return confirm('删除动作不可逆，请确定是否继续')" class="btn btn-outline-danger btn-sm">删除</button>
            <a href="{{ RouteName2URL "articles.edit" "id" .Article.GetStringID }}" class="btn btn-outline-secondary btn-sm">编辑</a>
        {{ end }}
        </form>

    </div><!-- /.blog-post -->
</div>

{{end}}