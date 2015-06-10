package frontend

import (
//	"fmt"
//	"encoding/json"
	"../../config"
	. "../../helpers/utilities"
	"../../models"
	//"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"gopkg.in/mgo.v2/bson"
)
//
//item
//
func (c Contrller) MypageItemGet(session sessions.Session,r render.Render,db DbSession) {
//	itemList := c.getItemTestData(session);
	
//	db.Write(itemList).Insert(itemList)
	//fmt.Println("errinsert:", errinsert)
	//bson.ObjectIdHex(id) //stringg to hex
	user := session.Get(config.SessionAuth)
	
	v, _ := user.(models.User)
	packlist := []models.Package{}
	db.Read(packlist).Find(bson.M{"owner": v.Id}).Sort("basemodel.date_insert").All(&packlist)
//	res, _ := json.Marshal(packlist)
//	fmt.Println(string(res))
	c.UserHTML(r, 200, "mypage/item", packlist)	
}
//
//item
//
func (c Contrller) MypageOrderGet(r render.Render) {
	c.UserHTML(r, 200, "mypage/order", nil)
}
