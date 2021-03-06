package requests

import (
    "github.com/thedevsaddam/govalidator"
    "goblog/app/models/category"
)

func ValidateCategoryForm(category category.Category) map[string][]string  {

    rules := govalidator.MapData{
        "name": []string{"required", "min_cn:1", "max_cn:10", "not_exists:categories,name"},
    }

    message := govalidator.MapData{
        "name" : []string{
            "required:标题为必填项",
            "min_cn:最短需要1个字",
            "max_cn:最长不能超过10个字",
        },
    }

    opts := govalidator.Options{
        Data:            &category,
        Rules:           rules,
        Messages:        message,
        TagIdentifier:   "valid",
    }

    return govalidator.New(opts).ValidateStruct()
}
