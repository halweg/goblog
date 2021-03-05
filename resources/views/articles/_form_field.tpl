
{{define "form-fields"}}
<div class="form-group mt-3">
    <label for="title">标题</label>
    <input type="text" class="form-control {{if .Errors.title }}is-invalid {{end}}" name="title" value="{{ .Title }}" required>
    {{ with .Errors.title }}
    <div class="invalid-feedback">
        {{ . }}
    </div>
    {{ end }}
</div>

<div class="form-group mt-3">
    <label for="body">内容</label>
    <textarea name="body" cols="30" rows="10" class="form-control {{if .Errors.body }}is-invalid {{end}}">{{ .Body }}</textarea>
    {{ with .Errors.body }}
    <div class="invalid-feedback">
        {{ . }}
    </div>
    {{ end }}
</div>
{{ end }}