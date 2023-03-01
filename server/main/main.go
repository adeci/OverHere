package main

import "OverHere/server/routes"

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

	router := routes.CreateRouter()
	routes.UserRoute(router)
	routes.ImageRoute(router)

	routes.Run(router)
}
