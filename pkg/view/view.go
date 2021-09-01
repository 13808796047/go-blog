package view

import (
	"github.com/13808796047/go-blog/app/models/category"
	"github.com/13808796047/go-blog/pkg/auth"
	"github.com/13808796047/go-blog/pkg/flash"
	"github.com/13808796047/go-blog/pkg/logger"
	"github.com/13808796047/go-blog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

// D 是 map[string]interface{} 的简写
type D map[string]interface{}

// Render 渲染通用视图
func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

// RenderSimple 渲染简单的视图
func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

// RenderTemplate 渲染视图
func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {
	var err error
	// 1. 通用模板数据
	data["isLogined"] = auth.Check()
	data["loginUser"] = auth.User
	data["flash"] = flash.All()
	data["Categories"], err = category.All()
	// 2. 生成模板文件
	allFiles := getTemplateFiles(tplFiles...)

	// 3. 解析所有模板文件
	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.RouteName2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)

	// 4. 渲染模板
	tmpl.ExecuteTemplate(w, name, data)
}

func getTemplateFiles(tplFiles ...string) []string {
	// 1.设置模板相对路径
	viewDir := "resources/views/"
	// 2. 遍历传参文件列表Slice,设置正确的路径,支持dir.filename语法糖
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}
	// 3.所有布局模板文件slice
	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)
	// 4.合并所有文件
	return append(layoutFiles, tplFiles...)
}
