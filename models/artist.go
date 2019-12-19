package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
	"time"
)

// Artist is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Artist struct {
	ID               uuid.UUID `json:"id" db:"id"`
	ArtistName       string    `json:"artist_name" db:"artist_name"`
	ArtistID         string    `json:"artist_id" db:"artist_id"`
	SpotifyURL       string    `json:"spotify_url" db:"spotify_url"`
	TimesRecommended int       `json:"times_recommended" db:"times_recommended"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (a Artist) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Artists is not required by pop and may be deleted
type Artists []Artist

// String is not required by pop and may be deleted
func (a Artists) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Artist) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: a.ArtistName, Name: "ArtistName"},
		&validators.StringIsPresent{Field: a.ArtistID, Name: "ArtistID"},
		&validators.StringIsPresent{Field: a.SpotifyURL, Name: "SpotifyURL"},
		&validators.IntIsPresent{Field: a.TimesRecommended, Name: "TimesRecommended"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Artist) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Artist) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
