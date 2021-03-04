{{define "app"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <title>{{template "title" .}}</title>
    <link href="/css/bootstrap.min.css" rel="stylesheet">
    <link href="/css/app.css" rel="stylesheet">
</head>

<body>

<div class="container-sm">
    <div class="row mt-5">

        {{template "sidebar" .}}

        {{template "main" .}}

    </div>
</div>

<script src="/js/bootstrap.min.js"></script>

</body>

</html>
{{end}}