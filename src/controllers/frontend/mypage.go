package frontend

import (
	//	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	. "../../helpers/utilities"
	//	"../../models"
	//	"gopkg.in/mgo.v2/bson"
)

func (c Contrller) MypageIndexGet(args martini.Params, r render.Render, db DbSession) {
	c.UserHTML(r, 200, "mypage/index", nil)
}
