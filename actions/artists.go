package actions

import (
	"net/http"
	"strings"

	"recommmended/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// ArtistsRecommendGet displays a form to add
// a new artists for future recommendations.
func ArtistsRecommendGet(c buffalo.Context) error {
	c.Set("artist", &models.Artist{})
	return c.Render(http.StatusOK, r.HTML("artists/recommend.html"))
}

// ArtistsRecommendPost first checks if the recommended artist
// has been already recommended. If not, then it makes
// a request to the Spotify API to get the artists info and if the artist
// does not exist it returns an error to the user. If exist it
// redirects the user to the home page with a flash message.
func ArtistsRecommendPost(c buffalo.Context) error {
	name := strings.ToLower(c.Params().Get("artist_name"))

	// Check if artist is already in DB
	var artist *models.Artist
	tx := c.Value("tx").(*pop.Connection)
	ok, err := tx.Where("artist_name = ?", name).Exists(artist)
	if err != nil {
		return errors.WithStack(err)
	}

	if ok {
		// TODO: Update artist
		//	TODO: Add mutexes
		c.Flash().Add("success", "Your recommendation has been readded successfully!")
		return c.Redirect(http.StatusFound, "/")
	}

	// Check in API
	artist, err = SpotifyClient.Search(name)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := tx.Create(artist); err != nil {
		return errors.WithStack(err)
	}

	// Success
	c.Flash().Add("success", "Your recommendation has been added successfully!")
	return c.Redirect(http.StatusFound, "/")
}
