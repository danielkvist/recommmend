package spotify

import (
	"context"
	"os"
	"recommmended/models"
	"strings"

	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// Client represents a client for the Spotify API.
type Client interface {
	SearchArtist(name string) (*models.Artist, error)
}

type client struct {
	c *spotify.Client
}

// New tries to create a valid Spotify client using
// environment variables for auth. If it fails to get
// a valid auth token it returns a non-nil error.
func New() (Client, error) {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		return nil, errors.Wrapf(err, "while getting authentication token")
	}

	c := spotify.Authenticator{}.NewClient(token)
	return &client{c: &c}, nil
}

// Search receives an artist's name and check if the artist
// exist in the Spotify database. If not exist or there is any
// problem while making the search it returns a non-nil error.
// If the artist exist it returns a new *model.Artist filled up.
func (c *client) SearchArtist(name string) (*models.Artist, error) {
	results, err := c.c.Search(name, spotify.SearchTypeArtist)
	if err != nil {
		return nil, errors.Wrapf(err, "while searching for artist %q", name)
	}

	if results.Artists == nil || len(results.Artists.Artists) == 0 {
		return nil, errors.Errorf("no results found searching for %q", name)
	}

	var am models.Artist
	artist := results.Artists.Artists[0]

	am.ArtistName = strings.ToLower(artist.SimpleArtist.Name)
	am.ArtistID = artist.ID.String()
	am.SpotifyURL = "https://open.spotify.com/artist/" + artist.ID.String()

	return &am, nil
}
