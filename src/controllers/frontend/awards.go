package frontend

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func (c Contrller) AwardsIndexGet(args martini.Params, r render.Render) {
	//r.JSON(200, r)
	for key, value := range args {
		print(key)
		print(value)
	}
	//render.Options
	opt := render.HTMLOptions{
		Layout: c.Options.Layout,
	}
	f := args["id"] + "topcontroller"
	r.HTML(200, c.ViewPath+"default/awards", f, opt)
}
