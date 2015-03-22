package utilities

import (
	"html/template"
	"strings"
	"time"
)

func GetFuncMap() template.FuncMap {
	f := funcMap{}
	return template.FuncMap{
		"nl2br":      f.nl2br,
		"htmlquote":  f.htmlquote,
		"str2html":   f.str2html,
		"dateformat": f.dateformat,
	}
}

type funcMap struct {
}

// 改行文字をbrタグに置き換える関数
func (f funcMap) nl2br(in string) (out string) {
	out = strings.Replace(in, "\n", "<br>", -1)
	return
}

// htmlquote
func (f funcMap) htmlquote(src string) string {
	text := string(src)
	text = strings.Replace(text, "&", "&amp;", -1)
	text = strings.Replace(text, "<", "&lt;", -1)
	text = strings.Replace(text, ">", "&gt;", -1)
	text = strings.Replace(text, "'", "&#39;", -1)
	text = strings.Replace(text, "\"", "&quot;", -1)
	text = strings.Replace(text, "“", "&ldquo;", -1)
	text = strings.Replace(text, "”", "&rdquo;", -1)
	text = strings.Replace(text, " ", "&nbsp;", -1)
	return strings.TrimSpace(text)
}

func (f funcMap) str2html(raw string) template.HTML {
	return template.HTML(raw)
}

func (f funcMap) dateformat(unixnano int64) string {
	return time.Unix(0, unixnano).Format("2006-01-02 15:04:05")
}
