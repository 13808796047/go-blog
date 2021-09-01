package controllers

import (
	"fmt"
	"github.com/13808796047/go-blog/app/models/article"
	"github.com/13808796047/go-blog/app/models/user"
	"github.com/13808796047/go-blog/pkg/logger"
	"github.com/13808796047/go-blog/pkg/route"
	"github.com/13808796047/go-blog/pkg/view"
	"net/http"
)

// UserController 用户控制器
type UserController struct {
	*BaseController
}

// Show 用户个人页面
func (u *UserController) Show(w http.ResponseWriter, r *http.Request) {

	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_user, err := user.Get(id)

	// 3. 如果出现错误
	if err != nil {
		u.ResponseForSQLError(w, err)
	} else {
		// ---  4. 读取成功，显示用户文章列表 ---
		articles, err := article.GetByUserID(_user.GetStringID())
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		} else {
			view.Render(w, view.D{
				"Articles": articles,
			}, "articles.index", "articles._article_meta")
		}
	}
}
