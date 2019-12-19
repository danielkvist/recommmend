package actions

import (
	"recommmended/models"

	"github.com/gobuffalo/buffalo"
)

// ArtistsRecommendGet displays a form to add
// a new artists for future recommendations.
func ArtistsRecommendGet(c buffalo.Context) error {
	c.Set("artist", &models.Artist{})
	return c.Render(200, r.HTML("artists/recommend.html"))
}
