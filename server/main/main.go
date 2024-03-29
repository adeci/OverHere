package main

import (
	"OverHere/server/routes"
	"OverHere/server/services/database"
)

func main() {
	// SPRINT 1
	// Demo workflow for uploading and retrieving image to and
	// database.DemoUploadAndRetrieveImage("images/spiderman.png")

	// Demo workflow for GIN setup.
	// handler.Handle()

	// SPRINT 2
	// Demo workflow for MongoDB data structure test
	// database.DemoDataStructureOHPostToImages("TEST1")
	// database.DemoDataStructureImagesToOHPost("username")
	// database.DemoDataStructureImagesToOHPost("username")
	// database.CreateAndStoreUserObject("ethan")

	database.ConnectMongoDBAtlas()
	router := routes.CreateRouter()
	routes.Route(router)
	routes.Run(router)
}
