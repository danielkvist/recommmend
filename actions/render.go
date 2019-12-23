package actions

import (
	"html/template"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
)

var r *render.Engine
var assetsBox *packr.Box

func init() {
	// Path changed for Google App Engine
	if ENV == "production" {
		assetsBox = packr.New("app:assets", "./public")
	} else {
		assetsBox = packr.New("app:assets", "../public")
	}

	r = render.New(render.Options{
		HTMLLayout:   "application.plush.html",
		TemplatesBox: packr.New("app:templates", "../templates"),
		AssetsBox:    assetsBox,

		Helpers: render.Helpers{
			"csrf": func() template.HTML {
				return template.HTML("<input name=\"authenticity_token\" value=\"<%= authenticity_token %>\" type=\"hidden\">")
			},
		},
	})
}
