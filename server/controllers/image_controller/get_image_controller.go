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

func GetImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		imageID := c.Param("imageid")
		defer cancel()

		fmt.Print("Getting image: " + imageID)

		retrievedImage, err := database.GetImage_ImageID(imageID)

		image := models.Image{
			ImageID:  retrievedImage.ImageID,
			UserID:   retrievedImage.UserID,
			OHPostID: retrievedImage.OHPostID,
			Encoding: retrievedImage.Base64Encode,
			XCoord:   retrievedImage.XCoord,
			YCoord:   retrievedImage.YCoord,
		}

		if err == nil {
			c.JSON(http.StatusOK, GetImageResponse(image))
		} else {
			c.JSON(http.StatusBadRequest, BadRequestImageResponse(err.Error()))
		}
	}
}

func GetImageResponse(retrievedImage models.Image) responses.ImageResponse {
	return responses.ImageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: map[string]interface{}{
			"data": retrievedImage,
		},
	}
}
