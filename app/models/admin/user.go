package admin

import (
    "goblog/app/models"
    "goblog/pkg/password"
)

// User 用户模型
type User struct {
    models.BaseModel

    UserName  string `gorm:"column:user_name;type:varchar(50);not null;unique" valid:"user_name"`
    Name     string `gorm:"column:name;type:varchar(255);not null;unique" valid:"name"`
    Email    string `gorm:"column:email;type:varchar(255);default:NULL;unique;" valid:"email"`
    Password string `gorm:"column:password;type:varchar(255)" valid:"password"`
    // gorm:"-" —— 设置 GORM 在读写时略过此字段
    PasswordConfirm string ` gorm:"-" valid:"password_confirm"`
}

func (User) TableName() string {
    //实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
    return "admin"
}

func (user User) ComparePassword(_password string) bool {
    return password.CheckHash(_password, user.Password)
}

func (user User) Link() string {
    //return route.Name2URL("users.show", "id", u.GetStringID())
    return ""
}
