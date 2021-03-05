{{define "title"}}
登录
{{end}}

{{define "main"}}
<div class="blog-post bg-white p-5 rounded shadow mb-4">

    <h3 class="mb-5 text-center">用户登录</h3>

    <form action="{{ RouteName2URL "auth.dologin" }}" method="post">

    <div class="form-group row mb-3">
        <label for="email" class="col-md-4 col-form-label text-md-right">E-mail</label>
        <div class="col-md-6">
            <input id="email" type="email" class="form-control {{if .Error }}is-invalid {{end}}" name="email" value="{{ .Email }}" required="">
            {{ with .Error }}
            <div class="invalid-feedback">
                <p>{{ . }}</p>
            </div>
            {{ end }}
        </div>
    </div>

    <div class="form-group row mb-3">
        <label for="password" class="col-md-4 col-form-label text-md-right">密码</label>
        <div class="col-md-6">
            <input id="password" type="password" class="form-control {{if .Errors.password }}is-invalid {{end}}" name="password" value="{{ .Password }}" required="">
        </div>
    </div>

    <div class="form-group row mb-3 mb-0 mt-4">
        <div class="col-md-6 offset-md-4">
            <button type="submit" class="btn btn-primary">
                登录
            </button>
        </div>
    </div>

    </form>

</div>


<div class="mb-3">
    <a href="/" class="text-sm text-muted"><small>返回首页</small></a>
    <a href="" class="text-sm text-muted float-right"><small>找回密码</small></a>
</div>

{{end}}