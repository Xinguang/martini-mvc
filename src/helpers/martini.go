package helpers

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/secure"
	. "github.com/starboychina/martini-mvc/src/helpers/utilities"
	"html/template"
	"io/ioutil"
)

type webConfig struct {
	Static    string   `json:"static"`
	Views     string   `json:"views"`
	Extension string   `json:"extension"`
	Admin     string   `json:"adminpath"`
	Database  DbConfig `json:"database"`
}

func Initialization() *martini.ClassicMartini {
	configfile, e := ioutil.ReadFile("config/config.json")
	var w webConfig = webConfig{
		Static:    "public",
		Views:     "views",
		Extension: ".tmpl",
		Admin:     "admin",
	}
	if e == nil {
		json.Unmarshal(configfile, &w)
	}
	return w.getMartini()
}

func (c *webConfig) getMartini() *martini.ClassicMartini {

	m := martini.Classic()
	//m.Use(martini.Logger())
	m.Use(martini.Recovery())
	if len(c.Static) != 0 {
		m.Use(martini.Static(c.Static))
	}
	m.Use(c.getRenderer())
	m.Use(DataHelper(c.Database))
	newRouter(m, c.Admin)
	m.Use(secure.Secure(secure.Options{
		SSLRedirect: true,
		SSLHost:     "localhost:443", // This is optional in production. The default behavior is to just redirect the request to the https protocol. Example: http://github.com/some_page would be redirected to https://github.com/some_page.
	}))
	//m.RunOnAddr(":80")
	//m.Run()
	return m
}

func (c *webConfig) getRenderer() martini.Handler {
	return render.Renderer(render.Options{
		Directory: c.Views, // Specify what path to load the templates from.
		//Layout:     "layout/layout",       // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions:      []string{c.Extension},            // Specify extensions to load for templates.
		Funcs:           []template.FuncMap{GetFuncMap()}, // Specify helper function maps for templates to access.
		Delims:          render.Delims{"{{", "}}"},        // Sets delimiters to the specified strings.
		Charset:         "UTF-8",                          // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON:      true,                             // Output human readable JSON
		IndentXML:       true,                             // Output human readable XML
		HTMLContentType: "text/html",                      // Output XHTML content type instead of default "text/html"
	})
}
