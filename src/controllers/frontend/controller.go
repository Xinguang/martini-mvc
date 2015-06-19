package frontend

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"../../config"
	. "../../helpers/utilities"
	
	"../../models"
	"gopkg.in/mgo.v2/bson"
)

type Contrller struct {
	Options
}

func (c Contrller) autoRedirect(session sessions.Session, r render.Render) {
	uri := session.Get(config.SessionRedirect)
	if nil == uri {
		r.Redirect("/", 301)
	} else {
		session.Delete(config.SessionRedirect)
		r.Redirect(uri.(string), 301)
	}
}


//
//item
//
func (c Contrller) getItemTestData(session sessions.Session) models.Package{
		
	itemList := models.Package{Id:bson.NewObjectId(),}
	
	user := session.Get(config.SessionAuth)
	if nil != user {
		v, ok := user.(models.User)
		if ok {
			itemList.Owner = v.Id 
		}
	}	
	Operation := GetOperation(session,"","")
	tracking := models.BoxTracking{Id:bson.NewObjectId(),};
	tracking.Shipping = models.Shipping{Id:bson.NewObjectId(),}
	tracking.Operation = Operation
	itemList.Tracking = []models.BoxTracking{tracking}
	
	boxsize := models.BoxSize{Id:bson.NewObjectId(),};
	boxsize.Operation = Operation
	itemList.Size = []models.BoxSize{boxsize}
	boximage := models.BoxImage{Id:bson.NewObjectId(),};
	boximage.Operation = Operation
	itemList.Image = []models.BoxImage{boximage}
	
	
	boxitem := models.Item{
		Id:bson.NewObjectId(),
		Tracking:[]models.BoxTracking{tracking},
		Size:[]models.BoxSize{boxsize},
		Image:[]models.BoxImage{boximage},
	};
	itemList.Item = []models.Item{boxitem}
	return itemList
}