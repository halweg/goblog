<!DOCTYPE html>
<html lang="en">
<head>
    <title>编辑文章 —— 我的技术博客</title>
    <style type="text/css">.error {color: red;}</style>
</head>
<body>
<form action="{{ .URL }}" method="post">
    <p><input type="text" name="title" value="{{ .Title }}"></p>
    {{ with .Errors.title }}
    <p class="error">{{ . }}</p>
    {{ end }}
    <p><textarea name="body" cols="30" rows="10">{{ .Body }}</textarea></p>
    {{ with .Errors.body }}
    <p class="error">{{ . }}</p>
    {{ end }}
    <p><button type="submit">更新</button></p>
</form>
</body>
</html>