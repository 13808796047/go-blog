package category

import (
	"github.com/13808796047/go-blog/pkg/logger"
	"github.com/13808796047/go-blog/pkg/model"
	"github.com/13808796047/go-blog/pkg/types"
)

// All 获取分类数据
func All() ([]Category, error) {
	var categories []Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

// Create 创建分类，通过 category.ID 来判断是否创建成功
func (category *Category) Create() (err error) {
	if err = model.DB.Create(&category).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// Get 通过 ID 获取分类
func Get(idstr string) (Category, error) {
	var category Category
	id := types.StringToInt(idstr)
	if err := model.DB.First(&category, id).Error; err != nil {
		return category, err
	}

	return category, nil
}
