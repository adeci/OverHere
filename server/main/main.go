package main

import "OverHere/server/routes/user_route"

func main() {
	// Demo workflow for uploading and retrieving image to and
	// from MongoDB
	//database.DemoUploadAndRetrieveImage("images/spiderman.png")

	router := user_route.CreateRouter()
	user_route.UserRoute(router)

	user_route.Run(router)
}
