package utilities

import (
	"github.com/martini-contrib/sessions"
	"gopkg.in/mgo.v2/bson"
	"../../models"
	"../../config"
	"time"
)

func GetOperation(session sessions.Session,Name string,Status string) models.Operation{
	
	auth := session.Get(config.SessionAuth)
	if nil != auth {
		user, ok := auth.(models.User)
		if ok {
			return models.Operation{
				Id:bson.NewObjectId(),
				Worker:user.Id,
				BaseModel:models.BaseModel{
					InsertDate:time.Now().Local(),
					UpdateDate:time.Now().Local(),
					//FlagDelete:false,
				},
			}
		}
	}	
	
	return models.Operation{}
}