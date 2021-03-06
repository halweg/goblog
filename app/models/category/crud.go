package category

import (
    "goblog/pkg/logger"
    "goblog/pkg/model"
    "goblog/pkg/types"
)

func (category *Category) Create() error {
    result := model.DB.Create(&category)
    if err := result.Error; err != nil {
        logger.LogError(err)
        return err
    }

    return nil
}

func All() ([]Category, error) {
    var categories []Category

    if err := model.DB.Find(&categories).Error; err != nil {
        return categories, err
    }
    return categories, nil
}

func Get(idstr string) (Category, error) {
    var category Category
    id := types.StringToInt(idstr)
    if err := model.DB.First(&category, id).Error; err != nil {
        return category, err
    }

    return category, nil
}
