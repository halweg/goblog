{{ define "user_profile_sticky_bar" }}
<style>
    .profile-bar-avatar { width: 48px; height: 48px; }
</style>
<div class="dropdown profile-element">
    <span>
        <img alt="image" class="img-circle profile-bar-avatar" src="{{ URL `/img/admin/default_avatar.png` }}  " />
    </span>
    <a data-toggle="dropdown" class="dropdown-toggle" href="#">
        <span class="clear">
            <span class="block m-t-xs"> <strong class="font-bold">Administrator</strong></span>
            <span class="text-muted text-xs block">超级管理员 <b class="caret"></b></span>
        </span>
    </a>
    <ul class="dropdown-menu animated fadeInRight m-t-xs">
        <li><a href="#">设置</a></li>
        <li><a href="#">联系人</a></li>
        <li><a href="#">邮箱</a></li>
        <li class="divider"></li>
        <li><a href="#">退出</a></li>
    </ul>
</div>
{{ end }}