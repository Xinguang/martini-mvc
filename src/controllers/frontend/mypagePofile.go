package frontend

import (
	//"fmt"
	//"encoding/json"
	//"../../config"
	. "../../helpers/utilities"
	//"../../models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	//"gopkg.in/mgo.v2/bson"
)

func (c Contrller) MypageIndexGet(r render.Render) {
	c.UserHTML(r, 200, "mypage/index", nil)
}
//
//Pofile
//
func (c Contrller) MypagePofileGet(session sessions.Session, r render.Render) {
	c.UserHTML(r, 200, "mypage/pofile", nil)
}

//
//Address
//
func (c Contrller) MypageAddressGet(session sessions.Session, r render.Render) {
	c.UserHTML(r, 200, "mypage/address", nil)
}
//
//Point
//
func (c Contrller) MypagePointGet(session sessions.Session, args martini.Params, r render.Render, db DbSession) {
	c.UserHTML(r, 200, "mypage/point", nil)
}
//
//Verified
//
func (c Contrller) MypageVerifiedGet(session sessions.Session, args martini.Params, r render.Render, db DbSession) {
	c.UserHTML(r, 200, "mypage/verified", nil)
}