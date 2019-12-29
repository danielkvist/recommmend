package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// SetErrorHandling adds custom error handling for certain errors
// as a 404 error or a 500 error.
func SetErrorHandling(app *buffalo.App) {
	app.ErrorHandlers[http.StatusNotFound] = errorHandler()
	app.ErrorHandlers[http.StatusInternalServerError] = errorHandler()
}

func errorHandler() buffalo.ErrorHandler {
	return func(status int, err error, c buffalo.Context) error {
		c.Logger().Error(err)
		c.Flash().Add("danger", fmt.Sprintf("You got a %v code! But don't worry, you're home now and we're aware of the problem.", status))
		if err := c.Redirect(http.StatusFound, "/"); err != nil {
			return err
		}

		return nil
	}
}
