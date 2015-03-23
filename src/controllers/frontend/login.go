package frontend

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/starboychina/martini-mvc/src/config"
)

func (c Contrller) LoginIndex(args martini.Params, r render.Render, session sessions.Session) {
	//r.JSON(200, r)
	v := session.Get(config.SessionRedirect)
	print("session:")
	if v != nil {
		print(v.(string))
	}

	for key, value := range args {
		print(key)
		print(value)
	}

	f := args["id"] + "topcontroller"
	r.HTML(200, c.ViewPath+"login/login", f)
}

func (c Contrller) LoginIndexPost(args martini.Params, r render.Render) {
	//r.JSON(200, r)
	for key, value := range args {
		print(key)
		print(value)
	}

	f := args["id"] + "topcontroller"
	r.HTML(200, c.ViewPath+"login/login", f)
}
