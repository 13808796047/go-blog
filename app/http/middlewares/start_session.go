package middlewares

import (
	"github.com/13808796047/go-blog/pkg/session"
	"net/http"
)

// StartSession开启session会话控制

func StartSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1.启动会话
		session.StartSession(w,r)
		// 2.继续处理接下去的请求
		next.ServeHTTP(w,r)
	})
}