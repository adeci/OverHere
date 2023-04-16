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

		_, err := database.GetOHPost_OHPostID(ohpostID)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestImageResponse(err.Error()),
			)
			return
		}

		databaseImage, err := database.GetImage_ImageID(imageID)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestImageResponse(err.Error()),
			)
			return
		}

		//Keep everything the same except OHPostId
		var databaseImageToPut database.ImageObject = database.ImageObject{
			ImageID:      databaseImage.ImageID,
			OHPostID:     ohpostID,
			UserID:       databaseImage.UserID,
			Base64Encode: databaseImage.Base64Encode,
			XCoord:       databaseImage.XCoord,
			YCoord:       databaseImage.YCoord,
		}

		database.PutImage(databaseImageToPut)

		retrievedImage, err := database.GetImage_ImageID(imageID)

		databaseImageChanged := models.Image{
			ImageID:  retrievedImage.ImageID,
			UserID:   retrievedImage.UserID,
			OHPostID: retrievedImage.OHPostID,
			Encoding: retrievedImage.Base64Encode,
			XCoord:   retrievedImage.XCoord,
			YCoord:   retrievedImage.YCoord,
		}

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
