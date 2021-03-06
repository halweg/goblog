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
