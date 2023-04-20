package image_controller

import (
	"OverHere/server/models"
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetImagesByTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		tag := c.Param("tag")
		defer cancel()

		fmt.Print("Getting image by userid: " + tag)

		allDatabaseImages, err := database.GetImage_All()

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestImageResponse(err.Error()),
			)
			cancel()
			return
		}

		matchingImages := []models.Image{}

		for _, databaseImage := range allDatabaseImages {
			if tag == databaseImage.Tag {
				image := models.Image{
					ImageID:  databaseImage.ImageID,
					OHPostID: databaseImage.OHPostID,
					UserID:   databaseImage.UserID,
					Encoding: databaseImage.Base64Encode,
					XCoord:   databaseImage.XCoord,
					YCoord:   databaseImage.YCoord,
					Tag:      databaseImage.Tag,
					Caption:  databaseImage.Caption,
				}

				matchingImages = append(matchingImages, image)
			}
		}

		if err == nil {
			c.JSON(http.StatusOK, GetMultipleImagesResponse(matchingImages))
		} else {
			c.JSON(http.StatusBadRequest, BadRequestImageResponse(err.Error()))
		}
	}
}
