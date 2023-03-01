package image_controller

import (
	"OverHere/server/models"
	"OverHere/server/responses"
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

		image := models.Image{
			ImageID:  imageID,
			OHPostID: "Test",
			Encoding: "Test",
		}

		c.JSON(http.StatusOK, GetImageResponse(image))
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
