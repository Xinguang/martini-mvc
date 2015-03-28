package frontend

import (
	//	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	. "github.com/starboychina/martini-mvc/src/helpers/utilities"
	//	"github.com/starboychina/martini-mvc/src/models"
	//	"gopkg.in/mgo.v2/bson"
)

func (c Contrller) AboutIndexGet(args martini.Params, r render.Render, db DbSession) {
	/*
		//r.JSON(200, r)
		for key, value := range args {
			print(key)
			print(value)
		}
		//render.Options
		opt := render.HTMLOptions{
			Layout: c.PathOptions.Layout,
		}
		f := args["id"] + "topcontroller"
		//////////////////////////////////////////
		user := models.User{}
		user.Id = bson.NewObjectId()
		user.Name = "real name"
		addr1 := models.Shipping{}
		addr1.Id = bson.NewObjectId()
		addr1.Name = "ship_name"
		user.Shippings = []models.Shipping{addr1}

		errinsert := db.Write(user).Insert(user)
		fmt.Println("errinsert:", errinsert)

		result := Person{}
		err := db.Read(result).Find(bson.M{"name": "Ale"}).One(&result)
		if err == nil {
			fmt.Println("Phone:", result.Phone)
		} else {
			fmt.Println("error:", err)
		}
		//////////////////////////////////////////

	*/
	c.HTML(r, 200, "default/about", nil)
}

type Person struct {
	Name  string
	Phone string
}
