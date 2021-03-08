{{ define "app" }}
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
</head>

<body>

<div id="wrapper">

    {{ template "slider" . }}

    <div id="page-wrapper" class="gray-bg">
        {{ template "nav" . }}

        {{ template "breadcrumb" . }}

        {{ template "main" . }}

        {{ template "footer" .}}

    </div>
</div>



<!-- Mainly scripts -->
<script src="/js/admin/jquery-2.1.1.js"></script>
<script src="/js/admin/bootstrap.min.js"></script>
<script src="/js/admin/plugins/metisMenu/jquery.metisMenu.js"></script>
<script src="/js/admin/plugins/slimscroll/jquery.slimscroll.min.js"></script>

<!-- Custom and plugin javascript -->
<script src="/js/admin/inspinia.js"></script>
<script src="/js/admin/plugins/pace/pace.min.js"></script>

</body>

</html>
{{ end }}
