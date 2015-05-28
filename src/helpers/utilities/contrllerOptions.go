package utilities

import (
	"github.com/martini-contrib/render"
	"../../models"
)

type Options struct {
	Layout   string
	ViewPath string
	User models.User
}

func (p Options) HTML(r render.Render, status int, name string, v interface{}) {
	opt := render.HTMLOptions{
		Layout: p.Layout,
	}
	r.HTML(status, p.ViewPath+name, v, opt)
}

func (p Options) UserHTML(r render.Render, status int, name string, v interface{}) {
	opt := render.HTMLOptions{
		Layout: "backend/layout/layout",
	}
	value := map[string]interface{}{
        "this":v,
        "session":p.User,
    };
	r.HTML(status, p.ViewPath+name, value, opt)
}
