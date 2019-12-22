package actions

import "net/http"

func (as *ActionSuite) Test_HomeHandler() {
	testGETHandler(as, "/", http.StatusOK)
}

func testGETHandler(as *ActionSuite, route string, expectedStatus int) {
	res := as.HTML(route).Get()
	as.Equal(expectedStatus, res.Code)
}
