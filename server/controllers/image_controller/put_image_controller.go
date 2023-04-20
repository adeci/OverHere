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

func PutImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		imageID := c.Param("imageid")
		defer cancel()

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

		fmt.Print("Getting image: " + imageID)

		var databaseImageToPut database.ImageObject = database.ImageObject{
			ImageID:      image.ImageID,
			OHPostID:     image.OHPostID,
			UserID:       image.UserID,
			Base64Encode: image.Encoding,
			XCoord:       image.XCoord,
			YCoord:       image.YCoord,
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
			Tag:      retrievedImage.Tag,
			Caption:  retrievedImage.Caption,
		}

		if err == nil {
			c.JSON(http.StatusOK, PutImageResponse(databaseImageChanged))
		} else {
			c.JSON(http.StatusBadRequest, BadRequestImageResponse(err.Error()))
		}
	}
}

func PutImageResponse(changedImage models.Image) responses.ImageResponse {
	return responses.ImageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: map[string]interface{}{
			"data": changedImage,
		},
	}
}
