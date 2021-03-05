package article

import (
	"goblog/pkg/model"
	"goblog/pkg/route"
)

type Article struct {

	model.BaseModel

    Title string
    Body string
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}




