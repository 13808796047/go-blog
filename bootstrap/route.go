package bootstrap

import (
	"github.com/13808796047/go-blog/pkg/route"
	"github.com/13808796047/go-blog/routes"
	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	route.SetRoute(router)
	return router
}
