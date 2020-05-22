
package main

import (
	"html/template"
	"github.com/unrolled/render"
)

var viewRender *render.Render

var noLayout = render.HTMLOptions{
	Layout: "",
}

func setupRender() {

	viewRender = render.New(render.Options{
		Layout: "layout",
		Extensions: []string{
			".html",
		},
		Funcs: []template.FuncMap{
			{
				"UnEscape": func(a string) template.HTML {
					return template.HTML(a)
				},
			},
		},
	})

}
