package frontend

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func (c Contrller) RegistIndex(args martini.Params, r render.Render) {
	//r.JSON(200, r)
	for key, value := range args {
		print(key)
		print(value)
	}

	f := args["id"] + "topcontroller"
	r.HTML(200, c.ViewPath+"user/login", f)
}
func (c Contrller) RegistIndexPost(args martini.Params, r render.Render) {
	//r.JSON(200, r)
	for key, value := range args {
		print(key)
		print(value)
	}

	f := args["id"] + "topcontroller"
	r.HTML(200, c.ViewPath+"user/login", f)
}
