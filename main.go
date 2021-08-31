package main

import (
	"github.com/13808796047/go-blog/app/http/middlewares"
	"github.com/13808796047/go-blog/bootstrap"
	"net/http"
)

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
