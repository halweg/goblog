package article

import (
	"goblog/app/models"
    "goblog/app/models/user"
    "goblog/pkg/route"
)

type Article struct {

	models.BaseModel
    UserID uint64 `gorm:"not null;index"`
    User   user.User
    Title string `valid:"title"`
    Body string `valid:"body"`

    CategoryID uint64 `gorm:"not null;default:3;index"`
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}

func (a Article) CreatedAtDate() string {
    return a.CreatedAt.Format("2006-01-02")
}




