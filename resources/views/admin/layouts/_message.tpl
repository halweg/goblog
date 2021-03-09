{{ define  "messages"}}

    {{ if .flash.danger }}
    <div class="alert alert-danger alert-dismissable">
        <button aria-hidden="true" data-dismiss="alert" class="close" type="button">×</button>
        {{ .flash.danger }}
    </div>
    {{ end }}

    {{ if .flash.warning }}
    <div class="alert alert-warning alert-dismissable">
        <button aria-hidden="true" data-dismiss="alert" class="close" type="button">×</button>
        {{ .flash.warning }}
    </div>
    {{ end }}

    {{ if .flash.success }}
    <div class="alert alert-success alert-dismissable">
        <button aria-hidden="true" data-dismiss="alert" class="close" type="button">×</button>
        {{ .flash.success }}
    </div>
    {{ end }}

    {{ if .flash.info }}
    <div class="alert alert-info alert-dismissable">
        <button aria-hidden="true" data-dismiss="alert" class="close" type="button">×</button>
        {{ .flash.info }}
    </div>
    {{ end }}

{{ end }}