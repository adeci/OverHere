package image_controller

import (
	"OverHere/server/controllers/helpers"
	"OverHere/server/models"
	"OverHere/server/responses"
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		fmt.Print("Creating image")

		var image models.Image

		//Validate the request body and bind
		if err := c.BindJSON(&image); err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestImageResponse(err.Error()),
			)
			return
		}

		if validationErr := helpers.Validate(&image); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestImageResponse(validationErr.Error()),
			)
			return
		}

		//Logic
		databaseImage, err := database.PostImage(image.Encoding, image.UserID, image.OHPostID, image.XCoord, image.YCoord, image.Tag, image.Caption)

		newImage := models.Image{
			ImageID:  databaseImage.ImageID,
			OHPostID: databaseImage.OHPostID,
			UserID:   databaseImage.UserID,
			Encoding: databaseImage.Base64Encode,
			XCoord:   databaseImage.XCoord,
			YCoord:   databaseImage.YCoord,
			Tag:      databaseImage.Tag,
			Caption:  databaseImage.Caption,
		}

		fmt.Print(newImage)

		if err == nil {
			//Successful response
			c.JSON(
				http.StatusCreated,
				PostImageResponse(newImage),
			)
		} else {
			c.JSON(http.StatusBadRequest, BadRequestImageResponse(err.Error()))
		}
	}
}

func PostImageResponse(newImage models.Image) responses.ImageResponse {
	return responses.ImageResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data: map[string]interface{}{
			"data": newImage,
		},
	}
}
