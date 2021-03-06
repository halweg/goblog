{{define "title"}}
创建文章分类
{{end}}

{{define "main"}}
<div class="col-md-9 blog-main">
    <div class="blog-post bg-white p-5 rounded shadow mb-4">

        <h3>新建文章分类</h3>

        <form action="{{ RouteName2URL "categories.store" }}" method="post">

        <div class="form-group mt-3">
            <label for="title">分类名称</label>
            <input type="text" class="form-control {{if .Errors.name }}is-invalid {{end}}" name="name" value="{{ .Category.Name }}" required>
            {{ with .Errors.name }}
            <div class="invalid-feedback">
                {{ . }}
            </div>
            {{ end }}
        </div>

        <button type="submit" class="btn btn-primary mt-3">提交</button>

        </form>

    </div><!-- /.blog-post -->
</div>

{{end}}