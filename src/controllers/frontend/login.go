package frontend

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/starboychina/martini-mvc/src/config"
	. "github.com/starboychina/martini-mvc/src/helpers/utilities"
	"github.com/starboychina/martini-mvc/src/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//ログイン
func (c Contrller) LoginIndex(req *http.Request, r render.Render, db DbSession, session sessions.Session) {
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
	c.HTML(r, 200, "login/login", nil)
}

//ログインチェック
func (c Contrller) LoginIndexPost(req *http.Request, r render.Render, db DbSession, session sessions.Session) {
	post := struct {
		Email    string
		Password string
		Remember string
		Message  string
	}{
		Email:    req.PostFormValue("email"),
		Password: req.PostFormValue("password"),
		Remember: req.PostFormValue("remember"),
		Message:  "",
	}

	if "1" == post.Remember {
		session.Options(sessions.Options{
			MaxAge: 3600 * 24 * 30, //30days
		})
	}
	user := models.User{}
	err := db.Read(user).Find(bson.M{"email": post.Email, "password": post.Password}).One(&user)
	if err == nil {
		session.Set(config.SessionAuth, user.Id.Hex())
		c.autoRedirect(session, r)
	} else {
		post.Message = "帐号或密码错误"
		c.HTML(r, 200, "login/login", post)
	}
}

//ログアウト
func (c Contrller) LogoutIndex(req *http.Request, r render.Render, db DbSession, session sessions.Session) {
	session.Delete(config.SessionAuth)
	c.autoRedirect(session, r)
}
