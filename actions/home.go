package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("index.html"))
}

// AboutHandler is a handler to serve up
// an about page.
func AboutHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("about.html"))
}
