package frontend

import (
	//"fmt"
	//"encoding/json"
	"../../config"
	. "../../helpers/utilities"
	"../../models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	//"gopkg.in/mgo.v2/bson"
)

func (c Contrller) MypageIndexGet(args martini.Params, r render.Render, db DbSession) {
	c.UserHTML(r, 200, "mypage/index", nil)
}
//
//Pofile
//
func (c Contrller) MypagePofileGet(session sessions.Session, args martini.Params, r render.Render, db DbSession) {
	user, ok := session.Get(config.SessionAuth).(models.User)
	if ok {
		c.UserHTML(r, 200, "mypage/pofile", user)
	}
}

//
//Address
//
func (c Contrller) MypageAddressGet(session sessions.Session, args martini.Params, r render.Render, db DbSession) {
	user, ok := session.Get(config.SessionAuth).(models.User)
	if ok {
		c.UserHTML(r, 200, "mypage/address", user)
	}
}
//
//Point
//
func (c Contrller) MypagePointGet(session sessions.Session, args martini.Params, r render.Render, db DbSession) {
	user, ok := session.Get(config.SessionAuth).(models.User)
	if ok {
		c.UserHTML(r, 200, "mypage/pofile", user)
	}
}
//
//Verified
//
func (c Contrller) MypageVerifiedGet(session sessions.Session, args martini.Params, r render.Render, db DbSession) {
	user, ok := session.Get(config.SessionAuth).(models.User)
	if ok {
		c.UserHTML(r, 200, "mypage/pofile", user)
	}
}