<!DOCTYPE html>
<html lang="en">
<head>
    <title>{{ .Title }} —— 我的技术博客</title>
    <style type="text/css">.error {color: red;}</style>
</head>
<body>
<p>ID： {{ .ID }}</p>
<p>标题： {{ .Title }}</p>
<p>内容：{{ .Body }}</p>

{{ $idString := Int64ToString .ID }}
<form action=" {{ RouteName2URL "articles.delete" "id" $idString  }} " method="POST">
    <button type="submit">删除</button>
</form>
</body>
</html>