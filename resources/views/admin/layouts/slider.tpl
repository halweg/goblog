{{ define "slider" }}
<nav class="navbar-default navbar-static-side" role="navigation">
    <div class="sidebar-collapse">
        <ul class="nav metismenu" id="side-menu">
            <li class="nav-header">
                {{ template "user_profile_sticky_bar" }}
                <div class="logo-element">
                </div>
            </li>
            <li class="active">
                <a href="#"><i class="fa fa-diamond"></i> <span class="nav-label">仪表盘</span></a>
            </li>
        </ul>

    </div>
</nav>
{{ end }}