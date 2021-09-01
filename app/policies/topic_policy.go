package policies

import (
	"github.com/13808796047/go-blog/app/models/article"
	"github.com/13808796047/go-blog/pkg/auth"
)

// CanModifyArticle 是否允许修改话题
func CanModifyArticle(_article *article.Article) bool {

	return auth.User().ID == _article.UserID
}
