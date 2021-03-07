{{define "title"}}
{{ .Title }}
{{end}}

{{define "main"}}

<div class="col-md-9 blog-main">
    <div class="blog-post bg-white p-5 rounded shadow mb-4">

        <div class="container">
            <div class="row">
                <div class="span10">
                    <div class="hero-unit">
                        <h1>{{ .Info }}</h1>
                        <p>{{ .Message }}</p>
                    </div>
                </div><!--/span-->
            </div>
        </div>

    </div><!-- /.blog-post -->
</div>

{{end}}