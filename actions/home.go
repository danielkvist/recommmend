package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// AboutHandler is a handler to serve up
// an about page.
func AboutHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("about.html"))
}
