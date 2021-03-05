{{define "title"}}
注册
{{end}}

{{define "main"}}
<div class="blog-post bg-white p-5 rounded shadow mb-4">

    <h3 class="mb-5 text-center">用户注册</h3>

    <form action="{{ RouteName2URL "auth.doregister" }}" method="post">

    <div class="form-group row mb-3">
        <label for="name" class="col-md-4 col-form-label text-md-right">姓名</label>
        <div class="col-md-6">
            <input id="name" type="text" class="form-control" name="name" value="" required="" autofocus="">
        </div>
    </div>

    <div class="form-group row mb-3">
        <label for="email" class="col-md-4 col-form-label text-md-right">E-mail</label>
        <div class="col-md-6">
            <input id="email" type="email" class="form-control" name="email" value="" required="">
        </div>
    </div>

    <div class="form-group row mb-3">
        <label for="password" class="col-md-4 col-form-label text-md-right">密码</label>
        <div class="col-md-6">
            <input id="password" type="password" class="form-control" name="password" required="">
        </div>
    </div>

    <div class="form-group row mb-3">
        <label for="password-confirm" class="col-md-4 col-form-label text-md-right">确认密码</label>
        <div class="col-md-6">
            <input id="password-confirm" type="password" class="form-control" name="password_confirmation" required="">
        </div>
    </div>

    <div class="form-group row mb-3 mb-0 mt-4">
        <div class="col-md-6 offset-md-4">
            <button type="submit" class="btn btn-primary">
                注册
            </button>
        </div>
    </div>

    </form>

</div>


<div class="mb-3">
    <a href="/" class="text-sm text-muted"><small>返回首页</small></a>
    <a href="/" class="text-sm text-muted float-right"><small>登录</small></a>
</div>

{{end}}