package article

import (
	"github.com/13808796047/go-blog/app/models"
	"github.com/13808796047/go-blog/pkg/route"
	"github.com/13808796047/go-blog/pkg/types"
)

// Article 文章模型
type Article struct {
	*models.BaseModel
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (a *Article) Link() string {
	return route.RouteName2URL("articles.show", "id", types.Uint64ToString(a.ID))
}
