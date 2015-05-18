package helpers

import (
	"encoding/json"
	"github.com/go-martini/martini"
	//"github.com/martini-contrib/csrf"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/secure"
	//"./config"
	. "./utilities"
	"html/template"
	"io/ioutil"
	//"net/http"
)

type webConfig struct {
	Static    string   `json:"static"`
	Views     string   `json:"views"`
	Extension string   `json:"extension"`
	Admin     string   `json:"adminpath"`
	Secret    string   `json:"secret"`
	Database  DbConfig `json:"database"`
}

func Initialization() *martini.Martini {
	configfile, e := ioutil.ReadFile("config/config.json")
	var w webConfig = webConfig{
		Static:    "public",
		Views:     "views",
		Extension: ".tmpl",
		Admin:     "admin",
		Secret:    "13B3AE09594A1A6B4C2E2046F098B6A6568E6F48979DB871A4FBC0FD861464FA", //demo.kansea.com
	}
	if e == nil {
		json.Unmarshal(configfile, &w)
	}
	return w.getMartini()
}

func (w *webConfig) getMartini() *martini.Martini {

	r := martini.NewRouter()
	newRouter(r, w.Admin)

	m := martini.New()
	//m.Use(martini.Logger())
	m.Use(martini.Recovery())
	if len(w.Static) != 0 {
		m.Use(martini.Static(w.Static, martini.StaticOptions{SkipLogging: true}))
	}
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	m.Use(w.getRenderer())
	m.Use(DataHelper(w.Database))
	m.Use(secure.Secure(secure.Options{
		SSLRedirect: true,
		SSLHost:     ":443", // This is optional in production. The default behavior is to just redirect the request to the https protocol. Example: http://github.com/some_page would be redirected to https://github.com/some_page.
	}))

	m.Use(w.Session())
	m.Use(w.Auth)

	return m

	/*
		m.Use(csrf.Generate(&csrf.Options{
			Secret:     w.Secret,
			SessionKey: config.SessionCsrf,
			// Custom error response.
			ErrorFunc: func(w http.ResponseWriter) {
				http.Error(w, "CSRF token validation failed", http.StatusBadRequest)
			}}))
	*/
	//m.RunOnAddr(":80")
	//m.Run()
}

func (w *webConfig) getRenderer() martini.Handler {
	return render.Renderer(render.Options{
		Directory: w.Views, // Specify what path to load the templates from.
		//Layout:     "layout/layout",       // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions:      []string{w.Extension},            // Specify extensions to load for templates.
		Funcs:           []template.FuncMap{GetFuncMap()}, // Specify helper function maps for templates to access.
		Delims:          render.Delims{"{{", "}}"},        // Sets delimiters to the specified strings.
		Charset:         "UTF-8",                          // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON:      true,                             // Output human readable JSON
		IndentXML:       true,                             // Output human readable XML
		HTMLContentType: "text/html",                      // Output XHTML content type instead of default "text/html"
	})
}
