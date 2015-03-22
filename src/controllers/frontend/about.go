package frontend

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	. "github.com/starboychina/martini-mvc/src/helpers/utilities"
	"gopkg.in/mgo.v2/bson"
)

func (c Contrller) AboutIndexGet(args martini.Params, r render.Render, db DbSession) {
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
	result := Person{}
	err := db.Read.C("people").Find(bson.M{"name": "Ale"}).One(&result)
	if err == nil {
		fmt.Println("Phone:", result.Phone)
	} else {
		fmt.Println("error:", err)
	}
	//////////////////////////////////////////

	r.HTML(200, c.ViewPath+"default/about", f, opt)
}

type Person struct {
	Name  string
	Phone string
}
