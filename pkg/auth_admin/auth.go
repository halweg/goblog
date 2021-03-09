package auth_admin
import (
    "errors"
    "goblog/app/models/admin"
    "goblog/pkg/session"
    "gorm.io/gorm"
)

func _getUID() string {
    _uid := session.Get("admin_uid")
    uid, ok := _uid.(string)
    if ok && len(uid) > 0 {
        return uid
    }
    return ""
}

// User 获取登录用户信息
func Admin() admin.User {
    uid := _getUID()
    if len(uid) > 0 {
        _user, err := admin.Get(uid)
        if err == nil {
            return _user
        }
    }
    return admin.User{}
}

func Attempt(username string, password string) error {
    _user, err := admin.GetByUserName(username)

    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("账号不存在或密码错误")
        } else {
            return errors.New("内部错误，请稍后尝试")
        }
    }

    if !_user.ComparePassword(password) {
        return errors.New("账号不存在或密码错误")
    }

    // 4. 登录用户，保存会话
    session.Put("admin_uid", _user.GetStringID())

    return nil
}

// Login 登录指定用户
func Login(_user admin.User) {
    session.Put("admin_uid", _user.GetStringID())
}

// Logout 退出用户
func Logout() {
    session.Forget("admin_uid")
}

// Check 检测是否登录
func Check() bool {
    return len(_getUID()) > 0
}
