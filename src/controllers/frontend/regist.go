package frontend

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	. "github.com/starboychina/martini-mvc/src/helpers/utilities"
	"github.com/starboychina/martini-mvc/src/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func (c Contrller) RegistIndex(r render.Render, db DbSession, session sessions.Session) {
	c.loginedAutoRedirect(session, r, db)
	c.HTML(r, 200, "regist/regist", nil)
}
func (c Contrller) RegistIndexPost(req *http.Request, r render.Render, db DbSession) {
	post := struct {
		User       models.User
		RePassword string
		Message    string
	}{
		User: models.User{
			Email:    req.PostFormValue("email"),
			Password: req.PostFormValue("password"),
			Permission:0
		},
		RePassword: req.PostFormValue("repassword"),
		Message:    "",
	}
	/*
		post := struct {
			models.User
			RePassword string
			Message    string
		}{
			models.User{
				Email:    req.PostFormValue("email"),
				Password: req.PostFormValue("password"),
			},
			req.PostFormValue("repassword"),
			"",
		}
	*/
	if len(post.User.Email) < 5 || len(post.User.Password) < 6 {
		post.Message = "请正确输入帐号密码"
		c.HTML(r, 403, "regist/regist", post)
		return
	}
	if post.User.Password != post.RePassword {
		post.Message = "密码不一致"
		c.HTML(r, 403, "regist/regist", post)
		return
	}
	err := db.Read(post.User).Find(bson.M{"email": post.User.Email}).One(&post.User)
	if nil == err {
		post.Message = "该Email已被注册过"
		post.User.Password = post.RePassword //防止密码泄漏
		c.HTML(r, 403, "regist/regist", post)
	} else {
		post.User.Id = bson.NewObjectId()
		aes := Aes{}
		post.User.Password = aes.AesEncrypt(post.User.Password)
		errinsert := db.Write(post.User).Insert(post.User)
		if nil == errinsert {
			r.Redirect("/login", 301)
		}
		r.Redirect("/503", 301)
	}
}
