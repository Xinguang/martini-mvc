package models

import (
	"gopkg.in/mgo.v2/bson"
)

//国际包裹
type Package struct { //打包收到的货 发往海外
	Id       bson.ObjectId `bson:"_id"`
	Owner    bson.ObjectId `bson:"owner"`    //所有者
	Tracking []BoxTracking `bson:"tracking"` //追跡(国际)
	Size     []BoxSize     `bson:"size"`     //箱子大小
	Image    []BoxImage    `bson:"image"`    //图片
	Note     string        `bson:"note"`     //备考
	Status   string        `bson:"status"`   //最新状态
	Item     []Item        `bson:"items"`    //货物列表
	BaseModel
}

//日本邮件
type Item struct { //日本国内收到的货
	Id       bson.ObjectId `bson:"_id"`
	Tracking []BoxTracking `bson:"tracking"` //追跡(日本国内)
	Size     []BoxSize     `bson:"size"`     //箱子大小
	Image    []BoxImage    `bson:"image"`    //图片
	Status   string        `bson:"status"`   //最新状态
	Note     string        `bson:"note"`     //备考
	BaseModel
}

//追跡
type BoxTracking struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`     //货名//邮单名
	Tracking string        `bson:"tracking"` //追跡番号
	Type     string        `bson:"type"`     //发送方式
	Postage  string        `bson:"postage"`  //邮费
	Shipping
	Operation
}

//箱子大小
type BoxSize struct {
	Id     bson.ObjectId `bson:"_id"`
	Weight int           `bson:"weight"` //重
	Width  int           `bson:"width"`  //宽
	Height int           `bson:"height"` //高
	Long   int           `bson:"long"`   //长
	Operation
}

//图片
type BoxImage struct {
	Id     bson.ObjectId `bson:"_id"`
	Name   string        `bson:"name"`   //图片名
	Width  string        `bson:"width"`  //宽
	Height string        `bson:"height"` //高
	Operation
}

//操作
type Operation struct {
	Id     bson.ObjectId `bson:"_id"`
	Name   string        `bson:"name"`   //货名
	Worker bson.ObjectId `bson:"worker"` //操作者
	Status string        `bson:"status"` //操作状态
	BaseModel
}

//寄送地址
type Shipping struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`     //收货人
	Zipcode  string        `bson:"zipcode"`  //收货邮编
	Address1 string        `bson:"address1"` //收货地址 省市
	Address2 string        `bson:"address2"` //收货地址 详细地址
	Tel      string        `bson:"tel"`      //收货人电话
	Mobile   string        `bson:"mobile"`   //收货人手机
}
