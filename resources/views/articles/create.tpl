{{define "title"}}
创建文章
{{end}}

{{define "main"}}
<div class="col-md-9 blog-main">
    <div class="blog-post bg-white p-5 rounded shadow mb-4">

        <h3>新建文章</h3>

        <form action="{{ RouteName2URL "articles.store" }}" method="post">

        {{ template "form-fields" . }}

        <button type="submit" class="btn btn-primary mt-3">提交</button>

        </form>

    </div><!-- /.blog-post -->
</div>

{{end}}