{{define "form-fields"}}
<div class="form-group mt-3">
    <label for="title">标题</label>
    <input id="title" type="text" class="form-control {{if .Errors.title }}is-invalid {{end}}" name="title" value="{{ .Article.Title }}" required>
    {{ with .Errors.title }}
    <div class="invalid-feedback">
        {{ . }}
    </div>
    {{ end }}
</div>

<div class="form-group mt-3">
    <label for="category">分类</label>
    <div class="form-group">
        <select class="form-control" id="category" name="category_id">
            {{ $articleCID := .Article.Category.ID }}
            {{ range $cat := .Category }}
                {{ if $articleCID }}
                    <option value="{{ $cat.ID }}" {{ if eq $cat.ID $articleCID }}selected="selected"{{ end }}> {{ $cat.Name }}  </option>
                    {{ else }}
                    <option value="{{ $cat.ID }}" > {{ $cat.Name }}  </option>
                {{ end }}
            {{ end }}
        </select>
    </div>
</div>

<div class="form-group mt-3">
    <label for="body">内容</label>
    <textarea id="body" name="body" cols="30" rows="10" class="form-control {{if .Errors.body }}is-invalid {{end}}">{{ .Article.Body }}</textarea>
    {{ with .Errors.body }}
    <div class="invalid-feedback">
        {{ . }}
    </div>
    {{ end }}
</div>
{{ end }}