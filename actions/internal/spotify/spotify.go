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

// Client represents a Spotify client.
type Client struct {
	client *spotify.Client
}

// New tries to create a valid Spotify client using
// environment variables for auth. If it fails to get
// a valid auth token it returns a non-nil error.
func New() (*Client, error) {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		return nil, errors.Wrapf(err, "while getting authentication token")
	}

	client := spotify.Authenticator{}.NewClient(token)
	return &Client{client: &client}, nil
}

func (c *Client) Search(artistName string) (*models.Artist, error) {
	results, err := c.client.Search(artistName, spotify.SearchTypeArtist)
	if err != nil {
		return nil, errors.Wrapf(err, "while searching for artist %q", artistName)
	}

	if results.Artists == nil || len(results.Artists.Artists) == 0 {
		return nil, errors.Errorf("no results found searching for %q", artistName)
	}

	var am models.Artist
	artist := results.Artists.Artists[0]

	am.ArtistName = strings.ToLower(artist.SimpleArtist.Name)
	am.ArtistID = artist.ID.String()
	am.SpotifyURL = "https://open.spotify.com/artist/" + artist.ID.String()

	return &am, nil
}
