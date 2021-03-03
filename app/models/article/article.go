package article

import (
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
    ID int64
    Title string
    Body string
}

func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(article.ID, 10))
}




