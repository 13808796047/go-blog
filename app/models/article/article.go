package article

import (
	"github.com/13808796047/go-blog/app/models"
	"github.com/13808796047/go-blog/app/models/user"
	"github.com/13808796047/go-blog/pkg/route"
	"github.com/13808796047/go-blog/pkg/types"
)

// Article 文章模型
type Article struct {
	*models.BaseModel
	Title  string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body   string `gorm:"type:text;not null;" valid:"body"`
	UserID uint64 `gorm:"not null;index"`
	User   *user.User
}

// Link 方法用来生成文章链接
func (a *Article) Link() string {
	return route.RouteName2URL("articles.show", "id", types.Uint64ToString(a.ID))
}

// CreatedAtDate 创建日期
func (a *Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}
