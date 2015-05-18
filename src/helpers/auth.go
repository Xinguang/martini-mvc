package helpers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"../config"
	. "./utilities"
	"net/http"
)

func (w webConfig) Session() martini.Handler {

	store := sessions.NewCookieStore([]byte(w.Secret))
	store.Options(sessions.Options{
		MaxAge: 0, //关闭浏览器 即失效
	})
	return sessions.Sessions("_session", store)
}
func (w webConfig) Auth(session sessions.Session, r render.Render, req *http.Request) {
	var reg ExRegexp = ExRegexp("/(mypage)|(" + w.Admin + ")/?")
	uri := req.URL.Path
	if reg.Match(uri) {
		v := session.Get(config.SessionAuth)
		if v == nil { //安全认证.......
			session.Set(config.SessionRedirect, uri)
			r.Redirect("/login", 302)
		} else { // 获取用户信息
		}
	}
	//session.Set("hello", "world")
}
