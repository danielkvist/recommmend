package actions

import "net/http"

func (as *ActionSuite) Test_ArtistsRecommendGet() {
	testGETHandler(as, "/artists/recommend", http.StatusOK)
}

func (as *ActionSuite) Test_ArtistsRecommendedGet() {
	testGETHandler(as, "/artists/recommended", http.StatusOK)
}
