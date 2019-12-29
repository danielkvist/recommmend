package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"

	"recommmended/actions/internal/spotify"
	"recommmended/models"

	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	csrf "github.com/gobuffalo/mw-csrf"
	i18n "github.com/gobuffalo/mw-i18n"
	"github.com/gobuffalo/packr/v2"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")

// T is used for translation.
var T *i18n.Translator

var app *buffalo.App
var spotifyClient spotify.Client

// App is where all routes and middleware for buffalo
// should be defined.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_recommmended_session",
		})

		// Tries to initialize a Spotify client.
		client, err := spotify.New()
		if err != nil {
			app.Stop(err)
		}
		spotifyClient = client

		// Adds custom error handling.
		SetErrorHandling(app)

		// Automatically redirect to SSL.
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		// Setup and use translations:
		app.Use(translations())

		// Main routes.
		app.GET("/", HomeHandler)
		app.GET("/about", AboutHandler)

		// Artists routes.
		artists := app.Group("/artists")
		artists.GET("/recommend", ArtistsRecommendGet)
		artists.POST("/recommend", ArtistsRecommendPost)
		artists.GET("/recommended", ArtistsRecommendedGet)

		// Serve files from the public directory
		app.ServeFiles("/", assetsBox)
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
func translations() buffalo.MiddlewareFunc {
	var err error

	path := "../locales"
	if ENV == "production" {
		path = "./locales"
	}

	if T, err = i18n.New(packr.New("app:locales", path), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS.
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
