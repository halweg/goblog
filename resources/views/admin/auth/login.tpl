{{define "title"}}
    登录
{{end}}

{{define "main"}}

<div class="row">

    <div class="col-md-6">
        <h2 class="font-bold">Welcome to GoBlog</h2>

        <p>
            GoBlog is build by go 1.15
        </p>

        <p>
            欢迎提出改进意见！
        </p>

        <p>
            <small>
                halweg@126.com
            </small>
        </p>

    </div>
    <div class="col-md-6">
        <div class="ibox-content">
            <form class="m-t" role="form" method="POST" action="{{ RouteName2URL "admin.auth.dologin" }}">
                <div class="form-group error">
                    <input type="text" class="form-control  {{if .Error }} error {{end}}   " placeholder="登录账户" required="" name="username">
                    {{ with .Error }}
                    <label  class="error">{{ . }}</label>
                    {{ end }}
                </div>
                <div class="form-group">
                    <input type="password" class="form-control" placeholder="密码" required="" name="password">
                </div>
                <button type="submit" class="btn btn-primary block full-width m-b">登录</button>

                <a href="#">
                    <small>忘记密码?</small>
                </a>

                <p class="text-muted text-center">
                    <small>还没有账号?</small>
                </p>
                <a class="btn btn-sm btn-white btn-block" href="#">注册</a>
            </form>
        </div>
    </div>
</div>
{{ end }}
