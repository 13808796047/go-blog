package article

import (
	"github.com/13808796047/go-blog/pkg/model"
	"github.com/13808796047/go-blog/pkg/types"
)

func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// Get通过ID获取文章
func Get(idstr string) (*Article, error) {
	var article *Article
	id := types.StringToInt(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}
