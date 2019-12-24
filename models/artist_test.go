package models

import "testing"

func Test_ArtistValidate(t *testing.T) {
	tt := []struct {
		name           string
		artist         Artist
		expectedToFail bool
	}{
		{
			name: "expected artist",
			artist: Artist{
				ArtistName:       "Jack White",
				ArtistID:         "notarealid",
				SpotifyURL:       "https://open.spotify.com/artist/4FZ3j1oH43e7cukCALsCwf",
				TimesRecommended: 1,
			},
		},
		{
			name: "artist without name",
			artist: Artist{
				ArtistID:         "notarealid",
				SpotifyURL:       "https://open.spotify.com/artist/4FZ3j1oH43e7cukCALsCwf",
				TimesRecommended: 0,
			},
			expectedToFail: true,
		},
		{
			name: "artist without id",
			artist: Artist{
				ArtistName:       "Jack White",
				SpotifyURL:       "https://open.spotify.com/artist/4FZ3j1oH43e7cukCALsCwf",
				TimesRecommended: 0,
			},
			expectedToFail: true,
		},
		{
			name: "artist without url",
			artist: Artist{
				ArtistName:       "Jack White",
				ArtistID:         "notarealid",
				TimesRecommended: 0,
			},
			expectedToFail: true,
		},
		{
			name: "expected artist",
			artist: Artist{
				ArtistName: "Jack White",
				ArtistID:   "notarealid",
				SpotifyURL: "https://open.spotify.com/artist/4FZ3j1oH43e7cukCALsCwf",
			},
			expectedToFail: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			valErrs, _ := tc.artist.Validate(nil)
			if valErrs.Count() != 0 {
				if tc.expectedToFail {
					t.Skipf("validation failed as expected: %v", valErrs)
				}

				t.Fatalf("while validating artist: %v", valErrs.String())
			}

			if tc.expectedToFail {
				t.Fatalf("validation expected to fail not failed")
			}
		})
	}
}
