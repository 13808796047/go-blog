package middlewares

import (
	"github.com/13808796047/go-blog/pkg/auth"
	"github.com/13808796047/go-blog/pkg/flash"
	"net/http"
)

// Auth登录用户才可以访问

func Auth(next HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !auth.Check() {
			flash.Warning("登录用户才能访问此页面")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next(w, r)
	}
}
