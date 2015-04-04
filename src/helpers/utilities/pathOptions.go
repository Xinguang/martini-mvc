package utilities

import (
	"github.com/martini-contrib/render"
)

type PathOptions struct {
	Layout   string
	ViewPath string
}

func (p PathOptions) HTML(r render.Render, status int, name string, v interface{}) {
	opt := render.HTMLOptions{
		Layout: p.Layout,
	}
	r.HTML(status, p.ViewPath+name, v, opt)
}

func (p PathOptions) UserHTML(r render.Render, status int, name string, v interface{}) {
	opt := render.HTMLOptions{
		Layout: "backend/layout/layout",
	}
	r.HTML(status, p.ViewPath+name, v, opt)
}
