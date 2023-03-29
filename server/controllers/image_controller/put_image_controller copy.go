package image_controller

import (
	"OverHere/server/responses"
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PutImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		imageID := c.Param("imageid")
		defer cancel()

		fmt.Print("Getting image: " + imageID)

		database.PutImage_XCoord(imageID, 80.0)
		//retrievedImage := database.PutImage_XCoord(imageID, 80.0)

		/*image := models.Image{
			ImageID:  retrievedImage.ImageID,
			UserID:   retrievedImage.UserID,
			OHPostID: retrievedImage.OHPostID,
			Encoding: retrievedImage.Base64Encode,
			XCoord:   retrievedImage.XCoord,
			YCoord:   retrievedImage.YCoord,
		}*/

		c.JSON(http.StatusOK, PutImageResponse())
	}
}

/*
func PutImageResponse(changedImage models.Image) responses.ImageResponse {
	return responses.ImageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: map[string]interface{}{
			"data": changedImage,
		},
	}
}
*/

func PutImageResponse() responses.ImageResponse {
	return responses.ImageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    map[string]interface{}{},
	}
}
