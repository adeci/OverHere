package image_controller

import (
	"OverHere/server/models"
	"OverHere/server/responses"
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PutAddImageToOHPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		imageID := c.Param("imageid")
		ohpostID := c.Param("ohpostid")

		defer cancel()

		fmt.Print("Getting image: " + imageID)

		oldOHPostImages, _ := database.GetImage_OHPostID(ohpostID)

		/*if err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestImageResponse(err.Error()),
			)
			cancel()
			return
		}*/

		//If already has image, don't put. Cancel.

		newDatabaseImage, err := database.GetImage_ImageID(imageID)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestImageResponse(err.Error()),
			)
			return
		}

		//Keep everything the same except OHPostId
		var databaseImageToPut database.ImageObject = database.ImageObject{
			ImageID:      newDatabaseImage.ImageID,
			OHPostID:     ohpostID,
			UserID:       newDatabaseImage.UserID,
			Base64Encode: newDatabaseImage.Base64Encode,
			XCoord:       newDatabaseImage.XCoord,
			YCoord:       newDatabaseImage.YCoord,
		}

		database.PutImage(databaseImageToPut)

		//Get image actually put into database. More of a test.
		retrievedImage, err := database.GetImage_ImageID(imageID)

		databaseImageChanged := models.Image{
			ImageID:  retrievedImage.ImageID,
			UserID:   retrievedImage.UserID,
			OHPostID: retrievedImage.OHPostID,
			Encoding: retrievedImage.Base64Encode,
			XCoord:   retrievedImage.XCoord,
			YCoord:   retrievedImage.YCoord,
		}

		oldOHPostImages = append(oldOHPostImages, retrievedImage)
		newAvgXCoord, newAvgYCoord := GetAverageCoordinatesFromImages(oldOHPostImages)
		database.PutOHPost_XCoord(ohpostID, newAvgXCoord)
		database.PutOHPost_YCoord(ohpostID, newAvgYCoord)

		if err == nil {
			c.JSON(http.StatusOK, PutAddImageToOHPostResponse(databaseImageChanged))
		} else {
			c.JSON(http.StatusBadRequest, BadRequestImageResponse(err.Error()))
		}
	}
}

func PutAddImageToOHPostResponse(changedImage models.Image) responses.ImageResponse {
	return responses.ImageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: map[string]interface{}{
			"data": changedImage,
		},
	}
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
