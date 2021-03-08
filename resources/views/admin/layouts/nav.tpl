{{ define "nav" }}
<div class="row border-bottom">
    <nav class="navbar navbar-static-top" role="navigation" style="margin-bottom: 0">
        <div class="navbar-header">
            <a class="navbar-minimalize minimalize-styl-2 btn btn-primary " href="#"><i class="fa fa-bars"></i> </a>
            <form role="search" class="navbar-form-custom" action="">
                <div class="form-group">
                    <input type="text" placeholder="" class="form-control" name="top-search" id="top-search">
                </div>
            </form>
        </div>
        <ul class="nav navbar-top-links navbar-right">

            <li class="dropdown">
                <a class="dropdown-toggle count-info" data-toggle="dropdown" href="#">
                    <i class="fa fa-envelope"></i>  <span class="label label-warning">开发中</span>
                </a>
                <ul class="dropdown-menu dropdown-messages">
                    <li>
                        <div class="dropdown-messages-box">
                            <a href="#" class="pull-left">
                                <img alt="image" class="img-circle" src=" {{ URL `/img/admin/default_avatar.png` }} ">
                            </a>
                            <div class="media-body">
                                <small class="pull-right">46小时前</small>
                                <strong>麦克先生</strong> 关注了 <strong>史密斯</strong>. <br>
                                <small class="text-muted">3 days ago at 7:58 pm - 10.06.2014</small>
                            </div>
                        </div>
                    </li>
                    <li class="divider"></li>
                    <li>
                        <div class="dropdown-messages-box">
                            <a href="#" class="pull-left">
                                <img alt="image" class="img-circle" src=" {{ URL `/img/admin/default_avatar.png` }}">
                            </a>
                            <div class="media-body ">
                                <small class="pull-right text-navy">5小时前</small>
                                <strong>克里斯</strong> 关注了 <strong>史密斯</strong>. <br>
                                <small class="text-muted">Yesterday 1:21 pm - 11.06.2014</small>
                            </div>
                        </div>
                    </li>
                    <li class="divider"></li>
                    <li>
                        <div class="dropdown-messages-box">
                            <a href="#" class="pull-left">
                                <img alt="image" class="img-circle" src=" {{ URL `/img/admin/default_avatar.png` }}">
                            </a>
                            <div class="media-body ">
                                <small class="pull-right">23小时前</small>
                                <strong>斯密斯</strong> 点赞了 <strong>吉米</strong>. <br>
                                <small class="text-muted">2 days ago at 2:30 am - 11.06.2014</small>
                            </div>
                        </div>
                    </li>
                    <li class="divider"></li>
                    <li>
                        <div class="text-center link-block">
                            <a href="#">
                                <i class="fa fa-envelope"></i> <strong>查看全部消息</strong>
                            </a>
                        </div>
                    </li>
                </ul>
            </li>
            <li class="dropdown">
                <a class="dropdown-toggle count-info" data-toggle="dropdown" href="#">
                    <i class="fa fa-bell"></i>  <span class="label label-primary">8</span>
                </a>
                <ul class="dropdown-menu dropdown-alerts">
                    <li>
                        <a href="#">
                            <div>
                                <i class="fa fa-envelope fa-fw"></i> 通知： 16条未读消息
                                <span class="pull-right text-muted small">4 分钟前</span>
                            </div>
                        </a>
                    </li>
                    <li class="divider"></li>
                    <li>
                        <a href="#">
                            <div>
                                <i class="fa fa-twitter fa-fw"></i>通知： 3 个新粉丝
                                <span class="pull-right text-muted small">12 分钟前</span>
                            </div>
                        </a>
                    </li>
                    <li class="divider"></li>
                    <li>
                        <a href="#">
                            <div>
                                <i class="fa fa-upload fa-fw"></i> 通知：服务重启了
                                <span class="pull-right text-muted small">4 分钟前</span>
                            </div>
                        </a>
                    </li>
                    <li class="divider"></li>
                    <li>
                        <div class="text-center link-block">
                            <a href="#">
                                <strong>查看全部提醒</strong>
                                <i class="fa fa-angle-right"></i>
                            </a>
                        </div>
                    </li>
                </ul>
            </li>


            <li>
                <a href="#">
                    <i class="fa fa-sign-out"></i> 登出
                </a>
            </li>
        </ul>

    </nav>
</div>
{{ end }}