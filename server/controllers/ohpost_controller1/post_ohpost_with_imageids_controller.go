package ohpost_controller1

//Followed tutorial for setup: https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m

import (
	"OverHere/server/controllers/helpers"
	"OverHere/server/models"
	"OverHere/server/services/database"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func PostOHPostWithImageIds() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var postOHPostRequest models.PostOHPostWithImageIds

		//Validate the request body
		if err := c.BindJSON(&postOHPostRequest); err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestOHPostResponse(err.Error()),
			)
			return
		}

		//Use the validator library to validate required fields
		if validationErr := helpers.Validate(&postOHPostRequest); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestOHPostResponse(validationErr.Error()),
			)
			cancel()
			return
		}

		//Logic
		//Averaging
		databaseOHPost, err := database.PostOHPost(postOHPostRequest.UserID, postOHPostRequest.Caption, 0.0, 0.0, postOHPostRequest.Tag)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestOHPostResponse(err.Error()),
			)
			cancel()
			return
		}

		//Update each image to have OHPost
		var databaseImages = []database.ImageObject{}

		for _, imageID := range postOHPostRequest.ImageIds {
			databaseImage, err := database.GetImage_ImageID(imageID)

			if err != nil {
				c.JSON(
					http.StatusBadRequest,
					BadRequestOHPostResponse(err.Error()),
				)
				cancel()
				return
			}

			databaseImages = append(databaseImages, databaseImage)
		}

		for _, databaseImage := range databaseImages {
			//Replace OHPost only
			var databaseImageToPut database.ImageObject = database.ImageObject{
				ImageID:      databaseImage.ImageID,
				OHPostID:     databaseOHPost.OHPostID,
				UserID:       databaseImage.UserID,
				Base64Encode: databaseImage.Base64Encode,
				XCoord:       databaseImage.XCoord,
				YCoord:       databaseImage.YCoord,
			}

			database.PutImage(databaseImageToPut)
		}

		avgXCoord, avgYCoord := GetAverageCoordinatesFromImages(databaseImages)

		//Successfully created, return result
		newOHPost := models.OHPost{
			OHPostID:  databaseOHPost.OHPostID,
			UserID:    databaseOHPost.UserID,
			Tag:       "Blank tag",
			Caption:   databaseOHPost.Description,
			AvgXCoord: avgXCoord,
			AvgYCoord: avgYCoord,
		}

		c.JSON(
			http.StatusCreated,
			PostedOHPostResponse(newOHPost),
		)
	}
}

func ToStringArray(str string) []string {
	return strings.Split(str, ",")
}

func GetAverageCoordinatesFromImages(images []database.ImageObject) (float64, float64) {
	sumXCoords := 0.0
	sumYCoords := 0.0
	totalImages := len(images)

	for _, image := range images {
		//Add to sum
		sumXCoords = sumXCoords + image.XCoord
		sumYCoords = sumYCoords + image.YCoord
	}

	avgXCoord := sumXCoords / float64(totalImages)
	avgYCoord := sumYCoords / float64(totalImages)

	return avgXCoord, avgYCoord
}
