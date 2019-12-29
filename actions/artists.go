package actions

import (
	"net/http"
	"strings"

	"recommmended/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
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
	tx := c.Value("tx").(*pop.Connection)
	ok, err := tx.Where("artist_name = ?", name).Exists(&models.Artist{})
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	if ok {
		err := tx.RawQuery("UPDATE artists SET times_recommended = times_recommended + 1 WHERE artist_name = ?", name).Exec()
		if err != nil {
			c.Logger().Error(err)
			return err
		}

		c.Flash().Add("success", "Your recommendation has been added successfully!")
		return c.Redirect(http.StatusFound, "/")
	}

	// Check on Spotify
	artist, err := spotifyClient.SearchArtist(name)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	if err := tx.Create(artist); err != nil {
		c.Logger().Error(err)
		return err
	}

	// Success
	c.Flash().Add("success", "Your recommendation has been added successfully!")
	return c.Redirect(http.StatusFound, "/")
}

// ArtistsRecommendedGet retrieves the last artist added
// to the database and shows it as a recommendation.
func ArtistsRecommendedGet(c buffalo.Context) error {
	var artist models.Artist
	tx := c.Value("tx").(*pop.Connection)

	// Get last artist added from the DB
	err := tx.Last(&artist)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	c.Set("artistName", strings.ToUpper(artist.ArtistName))
	c.Set("artistSpotifyURL", artist.SpotifyURL)
	return c.Render(http.StatusOK, r.HTML("artists/recommended.html"))
}
