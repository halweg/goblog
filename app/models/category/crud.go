package category

import (
    "goblog/pkg/logger"
    "goblog/pkg/model"
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
