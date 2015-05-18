package frontend

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"../../config"
	. "../../helpers/utilities"
	"../../models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//ログイン
func (c Contrller) LoginIndex(r render.Render, db DbSession, session sessions.Session) {
	c.loginedAutoRedirect(session, r, db)
	c.HTML(r, 200, "login/login", nil)
}

//ログインチェック
func (c Contrller) LoginIndexPost(req *http.Request, r render.Render, db DbSession, session sessions.Session) {

	post := struct {
		User     models.User
		Remember string
		Message  string
	}{
		User: models.User{
			Email:    req.PostFormValue("email"),
			Password: req.PostFormValue("password"),
		},
		Remember: req.PostFormValue("remember"),
		Message:  "",
	}

	if len(post.User.Email) < 5 || len(post.User.Password) < 6 {
		post.Message = "请正确输入帐号密码"
		c.HTML(r, 403, "login/login", post)
		return
	}
	if "1" == post.Remember {
		session.Options(sessions.Options{
			MaxAge: 3600 * 24 * 30, //30days
		})
	}
	aes := Aes{}
	post.User.Password = aes.AesEncrypt(post.User.Password)
	err := db.Read(post.User).Find(bson.M{"email": post.User.Email, "password": post.User.Password}).One(&post.User)
	if err == nil {
		session.Set(config.SessionAuth, post.User.Id.Hex())
		c.autoRedirect(session, r)
	} else {
		post.Message = "帐号或密码错误"
		post.User.Password = ""
		c.HTML(r, 403, "login/login", post)
	}
}

//ログアウト
func (c Contrller) LogoutIndex(req *http.Request, r render.Render, db DbSession, session sessions.Session) {
	session.Delete(config.SessionAuth)
	c.autoRedirect(session, r)
}

//
func (c Contrller) loginedAutoRedirect(session sessions.Session, r render.Render, db DbSession) {
	userid := session.Get(config.SessionAuth)
	//var user models.User
	if nil != userid {
		uid, ok := userid.(string)
		if ok {
			user := models.User{}
			err := db.Read(user).FindId(bson.ObjectIdHex(uid)).One(&user)
			if err == nil {
				c.autoRedirect(session, r)
			}
		}
	}
}
