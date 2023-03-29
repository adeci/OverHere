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

		err := database.DeleteImage_ImageID(imageID)

		if err == nil {
			c.JSON(http.StatusOK, DeleteImageResponse())
		} else {
			c.JSON(http.StatusBadRequest, BadRequestImageResponse(err.Error()))
		}
	}
}

func DeleteImageResponse() responses.ImageResponse {
	return responses.ImageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    map[string]interface{}{},
	}
}
