package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId `bson:"_id"`
	//Username  string        `bson:"username"` //用户名
	Email      string     `bson:"email"`      //Email
	Password   string     `bson:"password"`   //密码
	Name       string     `bson:"name"`       //姓名
	Note       string     `bson:"note"`       //备考
	Address    string     `bson:"address"`    //住址
	Tel        string     `bson:"tel"`        //电话
	Fax        string     `bson:"fax"`        //传真
	Mobile     string     `bson:"mobile"`     //手机
	QQ         string     `bson:"qq"`         //QQ
	Wechat     string     `bson:"wechat"`     //微信
	Point      int        `bson:"point"`      //积分
	Permission int        `bson:"permission"` //权限
	Shippings  []Shipping `bson:"shipping"`   //常用寄送地址
	BaseModel
}
