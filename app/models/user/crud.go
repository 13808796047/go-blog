package user

import (
	"github.com/13808796047/go-blog/pkg/logger"
	"github.com/13808796047/go-blog/pkg/model"
	"github.com/13808796047/go-blog/pkg/types"
)

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}
func Get(idstr string) (User, error) {
	var user User
	id := types.StringToInt(idstr)

	if err := model.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}
func GetByEmail(idstr string) (*User, error) {
	var user *User
	if err := model.DB.Where("email=?", idstr).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
