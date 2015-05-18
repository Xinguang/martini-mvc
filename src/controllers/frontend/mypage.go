package frontend

import (
	//	"fmt"
	//"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/starboychina/martini-mvc/src/config"
	. "github.com/starboychina/martini-mvc/src/helpers/utilities"
	"github.com/starboychina/martini-mvc/src/models"
	"gopkg.in/mgo.v2/bson"
)

func (c Contrller) MypageIndexGet(args martini.Params, r render.Render, db DbSession) {
	c.UserHTML(r, 200, "mypage/index", nil)
}

func (c Contrller) MypagePofileGet(session sessions.Session, args martini.Params, r render.Render, db DbSession) {
	uid, ok := session.Get(config.SessionAuth).(string)
	if ok {
		user := models.User{}
		err := db.Read(user).FindId(bson.ObjectIdHex(uid)).One(&user)

		//		b, err := json.Marshal(user)
		//		print(string(b))
		//{
		//    "Id": "5516d9f5887eb61b58000001",
		//    "Email": "admin@kansea.com",
		//    "Password": "8rduGlMgaKC1mInnbzRRgw==",
		//    "Name": "",
		//    "Note": "",
		//    "Address": "",
		//    "Tel": "",
		//    "Fax": "",
		//    "Mobile": "",
		//    "QQ": "",
		//    "Wechat": "",
		//    "Point": 0,
		//    "Permission": 0,
		//    "Shippings": [],
		//    "InsertDate": 0,
		//    "UpdateDate": 0,
		//    "FlagDelete": false
		//}
		if err == nil {
			c.UserHTML(r, 200, "mypage/pofile", user)
		}
	}
}
