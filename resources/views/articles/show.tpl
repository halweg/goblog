{{ define "title" }}
    {{ .Title }}
{{ end }}

{{ define "main" }}

<div class="col-md-9 blog-main">

    <div class="blog-post bg-white p-5 rounded shadow mb-4">
        <h3 class="blog-post-title">{{ .Title }}</h3>
        <p class="blog-post-meta text-secondary">发布于 <a href="" class="font-weight-bold">2020-09-05</a> by <a href="#" class="font-weight-bold">halweg</a></p>

        <hr>
        {{ .Body }}

        <form class="mt-4" action="{{ RouteName2URL "articles.delete" "id" .GetStringID }}" method="post">
        <button type="submit" onclick="return confirm('删除动作不可逆，请确定是否继续')">删除</button>
        </form>

    </div><!-- /.blog-post -->
</div>


{{ end }}