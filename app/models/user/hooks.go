package user

import (
	"github.com/13808796047/go-blog/pkg/password"
	"gorm.io/gorm"
)

// BeforeSave GORM 的模型钩子，在保存和更新模型前调用
func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	if !password.IsHashed(u.Password) {
		u.Password = password.Hash(u.Password)
	}
	return
}
