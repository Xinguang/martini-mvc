package helpers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	b "github.com/starboychina/martini-mvc/src/controllers/backend"
	f "github.com/starboychina/martini-mvc/src/controllers/frontend"
	. "github.com/starboychina/martini-mvc/src/helpers/utilities"
	"reflect"
	"strings"
)

//被映射 控制器 方法名 的规则
const _methodname ExRegexp = "(?P<controller>[A-Z][a-z]+)(?P<action>[A-Z][a-z]+)?(?P<method>[A-Z][a-z]+)?"
const _id string = "(:id)?"
const _slash string = "(\\/*)"

type appRouter struct {
	martini.Router
}

//创建路由
func newRouter(m martini.Router, adminpath string) {
	app := appRouter{m}
	app.backendRouter(adminpath)
	app.frontendRouter()
	app.NotFound(func(r render.Render) {
		opt := render.HTMLOptions{
			Layout: "frontend/layout/layout",
		}
		r.HTML(404, "404", "notfound", opt)
	})
}

//前台路由
func (app appRouter) frontendRouter() {
	base := PathOptions{Layout: "frontend/layout/layout", ViewPath: "frontend/"}
	c := f.Contrller{base}
	app.autoRouter(&c, _slash)
}

//后台路由
func (app appRouter) backendRouter(adminpath string) {
	base := PathOptions{Layout: "backend/layout/layout", ViewPath: "backend/"}
	c := b.Contrller{base}
	app.autoRouter(&c, _slash+adminpath+_slash)
}

//自动映射
func (app appRouter) autoRouter(i interface{}, groupurl string) {
	s := reflect.ValueOf(i).Elem()
	t := s.Type()
	app.Group(groupurl, func(r martini.Router) {
		for i := 0; i < s.NumMethod(); i++ {
			f := s.Method(i)
			methodName := t.Method(i).Name
			res := _methodname.FindAll(methodName)
			controller := strings.ToLower(res["controller"])
			action := strings.ToLower(res["action"])
			method := strings.ToLower(res["method"])
			if "default" == controller {
				r.Get(_slash, f.Interface())
				print(groupurl + _slash + "\n")
			} else if len(controller) > 0 && len(action) > 0 {
				url := controller + _slash + action + "(" + _slash + _id + ")?"
				app.setRouter(r, url, method, f.Interface())
				if "index" == action {
					app.setRouter(r, controller, method, f.Interface())
				}
			}
		}
	})
}

//根据请求方式设置接口
func (app appRouter) setRouter(r martini.Router, url string, method string, i interface{}) {
	switch method {
	case "post":
		r.Post(url, i)
	case "delete":
		r.Delete(url, i)
	case "patch":
		r.Patch(url, i)
	case "put":
		r.Put(url, i)
	default: //"get"
		r.Get(url, i)
	}
}
