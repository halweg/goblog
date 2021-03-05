package article

import (
    "goblog/pkg/logger"
    "goblog/pkg/model"
    "goblog/pkg/types"
)

func Get(idstr string) (Article, error) {
    var article Article
    id := types.StringToInt(idstr)
    if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
        return article, err
    }

    return article, nil
}

// GetAll 获取全部文章
func GetAll() ([]Article, error) {
    var articles []Article
    if err := model.DB.Debug().Preload("User").Find(&articles).Error; err != nil {
        return articles, err
    }
    return articles, nil
}

func (article *Article) Create() error {
	result := model.DB.Create(&article)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func (article *Article) Update() (int64,  error) {
	result := model.DB.Save(&article)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return result.RowsAffected, err
	}

	return result.RowsAffected,nil
}

func (article *Article) Delete() (int64, error) {

	result := model.DB.Delete(&article)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected,nil
}

func GetByUserID(uid string) ([]Article, error) {
    var articles []Article
    if err := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Error; err != nil {
        return articles, err
    }

    return articles, nil
}