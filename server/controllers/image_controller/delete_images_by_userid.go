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

func DeleteImagesByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userID := c.Param("userid")
		defer cancel()

		fmt.Print("Deleting image by UserID: " + userID)

		err := database.DeleteImage_UserID(userID)

		if err == nil {
			c.JSON(http.StatusOK, DeleteMultipleImagesResponse())
		} else {
			c.JSON(http.StatusBadRequest, BadRequestImageResponse(err.Error()))
		}
	}
}

func DeleteMultipleImagesResponse() responses.ImageResponse {
	return responses.ImageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    map[string]interface{}{},
	}
}
