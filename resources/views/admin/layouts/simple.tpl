{{define "simple"}}
<!DOCTYPE html>
<html>

<head>

    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>{{template "title" .}}</title>

    <link href="/css/admin/bootstrap.min.css" rel="stylesheet">
    <link href="/css/admin/font-awesome/css/font-awesome.css" rel="stylesheet">

    <link href="/css/admin/animate.css" rel="stylesheet">
    <link href="/css/admin/style.css" rel="stylesheet">
    <style>
        .main {
            max-width: 800px;
            margin: 0 auto;
            padding: 100px 20px 20px 20px;
        }
    </style>
</head>

<body class="gray-bg">

<div class="main animated fadeInDown">
    {{template "main" .}}
    <hr/>
    <div class="row">
        <div class="col-md-6">
            Copyright GoBlog
        </div>
        <div class="col-md-6 text-right">
            <small>© 2020-2021</small>
        </div>
    </div>
</div>

</body>

</html>
{{ end }}
