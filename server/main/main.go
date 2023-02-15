package main

import "OverHere/server/database"

func main() {
	// Demo workflow for uploading and retrieving image to and
	// from MongoDB
	database.DemoUploadAndRetrieveImage("images/spiderman.png")
}
