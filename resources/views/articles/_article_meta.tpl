{{define "article-meta"}}
<p class="blog-post-meta text-secondary">
    发布于 <a href="{{ .Link }}" class="font-weight-bold">{{ .CreatedAtDate }}</a>
    by <a href="{{ .User.Link }}" class="font-weight-bold">{{ .User.Name }}</a>
</p>
{{ end }}