package models

import (
	"time"
)

type BaseModel struct {
	InsertDate time.Time `bson:"date_insert"` //写入时间
	UpdateDate time.Time `bson:"date_update"` //更新时间
	FlagDelete bool                `bson:"flag_delete"` //可用标识
}
