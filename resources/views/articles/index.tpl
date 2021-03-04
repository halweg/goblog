<!DOCTYPE html>
<html lang="en">

<head>
    <title>所有文章 —— 我的技术博客</title>
    <link href="/css/bootstrap.min.css" rel="stylesheet">
    <link href="/css/app.css" rel="stylesheet">
</head>

<body>

<div class="container-sm">
    <div class="row mt-5">

        <div class="col-md-3 blog-sidebar">
            <div class="p-4 mb-3 bg-white rounded shadow-sm">
                <h1>GoBlog</h1>
                <p class="mb-0">摒弃世俗浮躁，追求技术精湛</p>
            </div>

            <div class="p-4 bg-white rounded shadow-sm mb-3">
                <h5>分类</h5>
                <ol class="list-unstyled mb-0">
                    <li><a href="#">未分类</a></li>
                    <li><a href="#">模板</a></li>
                    <li><a href="#">数据库</a></li>
                </ol>
            </div>

            <div class="p-4 bg-white rounded shadow-sm mb-3">
                <h5>作者</h5>
                <ol class="list-unstyled mb-0">
                    <li><a href="#">Summer</a></li>
                    <li><a href="#">Aufree</a></li>
                    <li><a href="#">Monkey</a></li>
                </ol>
            </div>

            <div class="p-4 bg-white rounded shadow-sm mb-3">
                <h5>链接</h5>
                <ol class="list-unstyled">
                    <li><a href="#">关于我们</a></li>
                    <li><a href="#">注册</a></li>
                    <li><a href="#">登录</a></li>
                </ol>
            </div>
        </div><!-- /.blog-sidebar -->

        <div class="col-md-9 blog-main">

            {{ range $key, $article := . }}

            <div class="blog-post bg-white p-5 rounded shadow mb-4">
                <h3 class="blog-post-title"><a href="{{ $article.Link }}" class="text-dark text-decoration-none">{{ $article.Title }}</a></h3>
                <p class="blog-post-meta text-secondary">发布于 <a href="" class="font-weight-bold">2020-09-05</a> by <a href="#" class="font-weight-bold">Summer</a></p>

                <hr>
                {{ $article.Body }}

            </div><!-- /.blog-post -->

            {{ end }}


            <nav class="blog-pagination mb-5">
                <a class="btn btn-outline-primary" href="#">下一页</a>
                <a class="btn btn-outline-secondary disabled" href="#" tabindex="-1" aria-disabled="true">上一页</a>
            </nav>

        </div><!-- /.blog-main -->

    </div>
</div>

<script src="/js/bootstrap.min.js"></script>

</body>

</html>