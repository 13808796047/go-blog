package route

import (
	"github.com/13808796047/go-blog/pkg/config"
	"github.com/13808796047/go-blog/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

var route *mux.Router

// SetRoute 设置路由实例,以供RouteName2URL等函数使用
func SetRoute(r *mux.Router) {
	route = r
}

// RouteName2URL 通过路由名称来获取URL
func RouteName2URL(routeName string, pairs ...string) string {

	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}

	return config.GetString("app.url") + url.String()
}

// GetRouteVariable 获取URI路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
