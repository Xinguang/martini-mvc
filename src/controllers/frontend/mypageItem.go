package frontend

import (
	//"fmt"
	//"encoding/json"
	//"../../config"
	//. "../../helpers/utilities"
	"../../models"
	//"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	//"gopkg.in/mgo.v2/bson"
)
//
//item
//
func (c Contrller) MypageItemGet(session sessions.Session, r render.Render) {
	itemList := models.Package{
	}
	c.UserHTML(r, 200, "mypage/item", itemList)
}
//
//item
//
func (c Contrller) MypageOrderGet(session sessions.Session, r render.Render) {
	c.UserHTML(r, 200, "mypage/order", nil)
}
