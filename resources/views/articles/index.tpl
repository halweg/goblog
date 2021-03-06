{{define "title"}}
所有文章 —— 我的技术博客
{{end}}

{{define "main"}}
<div class="col-md-9 blog-main">

    {{ range $key, $article := .Articles }}

    <div class="blog-post bg-white p-5 rounded shadow mb-4">
        <h3 class="blog-post-title"><a href="{{ $article.Link }}" class="text-dark text-decoration-none">{{ $article.Title }}</a></h3>
        {{template "article-meta" $article }}

        <hr>
        {{ $article.Body }}

    </div><!-- /.blog-post -->

    {{ end }}

    {{template "pagination" .PagerData }}

</div><!-- /.blog-main -->
{{end}}