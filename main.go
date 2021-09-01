package main

import (
	"github.com/13808796047/go-blog/app/http/middlewares"
	"github.com/13808796047/go-blog/bootstrap"
	"github.com/13808796047/go-blog/config"
	c "github.com/13808796047/go-blog/pkg/config"
	"net/http"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
