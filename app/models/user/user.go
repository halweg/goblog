package user

import "goblog/app/models"

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"column:email;type:varchar(255);default:NULL;unique;" valid:"email"`
	Password string `gorm:"column:password;type:varchar(255)" valid:"password"`
	// gorm:"-" —— 设置 GORM 在读写时略过此字段
	PasswordConfirm string ` gorm:"-" valid:"password_confirm"`
}

func (u User) ComparePassword(password string) bool  {
	return u.Password == password
}