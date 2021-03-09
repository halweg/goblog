{{ define  "messages"}}

    {{ if .flash.danger }}
        <div class="flash-message">
            <div class="alert alert-danger">
                {{ .flash.danger }}
            </div>
        </div>
    {{ end }}

    {{ if .flash.warning }}
        <div class="flash-message">
            <div class="alert alert-warning">
                {{ .flash.warning }}
            </div>
        </div>
    {{ end }}

    {{ if .flash.success }}
        <div class="flash-success">
            <div class="alert alert-success">
                {{ .flash.success }}
            </div>
        </div>
    {{ end }}

    {{ if .flash.info }}
        <div class="flash-info">
            <div class="alert alert-info">
                {{ .flash.info }}
            </div>
        </div>
    {{ end }}

{{ end }}

