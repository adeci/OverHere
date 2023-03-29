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

func DeleteImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		imageID := c.Param("imageid")
		defer cancel()

		fmt.Print("Getting image: " + imageID)

		database.DeleteImage_ImageID(imageID)

		c.JSON(http.StatusOK, DeleteImageResponse())
	}
}

func DeleteImageResponse() responses.ImageResponse {
	return responses.ImageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    map[string]interface{}{},
	}
}
