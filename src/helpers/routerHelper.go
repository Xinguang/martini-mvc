package helpers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	b "github.com/starboychina/martini-mvc/src/controllers/backend"
	f "github.com/starboychina/martini-mvc/src/controllers/frontend"
	. "github.com/starboychina/martini-mvc/src/options"
	"reflect"
	"strings"
)

const _methodname ExRegexp = "(?P<controller>[A-Z][a-z]+)(?P<action>[A-Z][a-z]+)(?P<method>[A-Z][a-z]+)"
const _id string = "(:id)?"
const _slash string = "(\\/+)"

type appRouter struct {
	*martini.ClassicMartini
}

func newRouter(m *martini.ClassicMartini, adminpath string) {
	app := appRouter{m}
	app.backendRouter(adminpath)
	app.frontendRouter()
	app.NotFound(func(r render.Render) {
		opt := render.HTMLOptions{
			Layout: "backend/layout/layout",
		}
		r.HTML(404, "404", "notfound", opt)
	})
}
func (app appRouter) frontendRouter() {
	base := PathOptions{Layout: "frontend/layout/layout", ViewPath: "frontend/"}
	c := f.Contrller{base}
	app.autoRouter(&c, _slash)
}

func (app appRouter) backendRouter(adminpath string) {
	base := PathOptions{Layout: "backend/layout/layout", ViewPath: "backend/"}
	c := b.Contrller{base}
	app.autoRouter(&c, _slash+adminpath+_slash)
}

func (app appRouter) autoRouter(i interface{}, groupurl string) {
	s := reflect.ValueOf(i).Elem()
	t := s.Type()
	app.Group(groupurl, func(r martini.Router) {
		for i := 0; i < s.NumMethod(); i++ {
			f := s.Method(i)
			methodName := t.Method(i).Name
			res := _methodname.findAll(methodName)
			controller := strings.ToLower(res["controller"])
			action := strings.ToLower(res["action"])
			method := strings.ToLower(res["method"])
			if len(controller) > 0 && len(action) > 0 {
				url := controller + _slash + action + "(" + _slash + _id + ")?"
				switch method {
				case "post":
					r.Post(url, f.Interface())
				case "delete":
					r.Delete(url, f.Interface())
				case "patch":
					r.Patch(url, f.Interface())
				case "put":
					r.Put(url, f.Interface())
				default: //"get"
					r.Get(url, f.Interface())
					if "index" == action {
						r.Get(controller, f.Interface())
					}
				}
			}
		}
	})
}
