package models

import (
	"gopkg.in/mgo.v2/bson"
)

type BaseModel struct {
	InsertDate bson.MongoTimestamp `bson:"date_insert"` //写入时间
	UpdateDate bson.MongoTimestamp `bson:"date_update"` //更新时间
	FlagDelete bool                `bson:"flag_delete"` //可用标识
}
