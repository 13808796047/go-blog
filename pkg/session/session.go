package session

import (
	"github.com/13808796047/go-blog/pkg/config"
	"github.com/13808796047/go-blog/pkg/logger"
	"github.com/gorilla/sessions"
	"net/http"
)

// Store gorilla session的存储库
var Store = sessions.NewCookieStore([]byte("33446a9dcf9ea060a0a6532b166da32f304af0de"))

// Session 当前会话
var Session *sessions.Session

// Request用于获取会话

var Request *http.Request

// Response 用于写入会话

var Response http.ResponseWriter

func StartSession(w http.ResponseWriter, r *http.Request) {
	var err error
	// Store.Get()的第二个参数是Cookie的名称
	// gorilla/session支持多会话,本项目我们只使用单一会话即可
	Session, err = Store.Get(r, config.GetString("session.session_name"))
	logger.LogError(err)
	Request = r
	Response = w
}

// Put写入键值对应的会话数据
func Put(key string, value interface{}) {
	Session.Values[key] = value
	Save()
}

// Get获取会话数据,获取数据时请做类型检测
func Get(key string) interface{} {
	return Session.Values[key]
}

// forget 删除某个会话项
func Forget(key string) {
	delete(Session.Values, key)
	Save()
}

// Flush 删除当前会话
func Flush() {
	Session.Options.MaxAge = -1
	Save()
}

// Save保持会话
func Save() {
	// 非HTTPS的链接无法使用Secure和HttpOnly,浏览器会报错
	// Session.Options.Secure = true
	// Session.Options.HttpOnly = true
	err := Session.Save(Request, Response)
	logger.LogError(err)
}
