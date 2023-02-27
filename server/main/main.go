package main

import "OverHere/server/routes/user_route"

func main() {
	// SPRINT 1
	// Demo workflow for uploading and retrieving image to and
	// from MongoDB
	// database.DemoUploadAndRetrieveImage("images/spiderman.png")

	// Demo workflow for GIN setup.
	// handler.Handle()

	// SPRINT 2
	// Demo workflow for logging in, creating post w/ images,
	// and displaying posts by a user. (MongoDB data structure test)
	//database.DemoDataStructure()
	//database.DemoUploadAndRetrieveImage("images/spiderman.png")

	router := user_route.CreateRouter()
	user_route.UserRoute(router)

	user_route.Run(router)
}
