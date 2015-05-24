package frontend

import (
	"fmt"
	//"encoding/json"
	"../../config"
	. "../../helpers/utilities"
	"../../models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"gopkg.in/mgo.v2/bson"
)

func (c Contrller) MypageIndexGet(args martini.Params, r render.Render, db DbSession) {
	c.UserHTML(r, 200, "mypage/index", nil)
}
//
//Pofile
//
func (c Contrller) MypagePofileGet(session sessions.Session, args martini.Params, r render.Render, db DbSession) {
	uid, _ := session.Get(config.SessionAuth).(string)
	user := models.User{}
	err := db.Read(user).FindId(bson.ObjectIdHex(uid)).One(&user)
	if err == nil {
		c.UserHTML(r, 200, "mypage/pofile", user)
	}
}

//
//Address
//
func (c Contrller) MypageAddressGet(session sessions.Session, args martini.Params, r render.Render, db DbSession) {
	uid, _ := session.Get(config.SessionAuth).(string)
	user := models.User{}
	err := db.Read(user).FindId(bson.ObjectIdHex(uid)).One(&user)
	if err == nil {
		fmt.Println(user.Shippings)
		c.UserHTML(r, 200, "mypage/address", user)
	}
}
//
//Point
//
func (c Contrller) MypagePointGet(session sessions.Session, args martini.Params, r render.Render, db DbSession) {
	uid, _ := session.Get(config.SessionAuth).(string)
	user := models.User{}
	err := db.Read(user).FindId(bson.ObjectIdHex(uid)).One(&user)
	if err == nil {
		c.UserHTML(r, 200, "mypage/pofile", user)
	}
}
//
//Verified
//
func (c Contrller) MypageVerifiedGet(session sessions.Session, args martini.Params, r render.Render, db DbSession) {
	uid, _ := session.Get(config.SessionAuth).(string)
	user := models.User{}
	err := db.Read(user).FindId(bson.ObjectIdHex(uid)).One(&user)
	if err == nil {
		c.UserHTML(r, 200, "mypage/pofile", user)
	}
}