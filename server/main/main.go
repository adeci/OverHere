package main

import "OverHere/server/services/database"

func main() {
	// SPRINT 1
	// Demo workflow for uploading and retrieving image to and
	// database.DemoUploadAndRetrieveImage("images/spiderman.png")

	// Demo workflow for GIN setup.
	// handler.Handle()

	// SPRINT 2
	// Demo workflow for MongoDB data structure test
	// database.DemoDataStructureOHPostToImages("user")
	// database.DemoDataStructureImagesToOHPost("username")
	database.CreateAndStoreUserObject("hello")

	// router := user_route.CreateRouter()
	// user_route.UserRoute(router)

	// user_route.Run(router)
}
