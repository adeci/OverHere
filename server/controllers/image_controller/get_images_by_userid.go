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

func GetImagesByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userID := c.Param("userid")
		defer cancel()

		fmt.Print("Getting image by userid: " + userID)

		allDatabaseImages, err := database.GetImage_All()

		if err != nil {
			fmt.Println(err)
			cancel()
			return
		}

		matchingImages := []models.Image{}

		for _, databaseImage := range allDatabaseImages {
			if userID == databaseImage.UserID {
				image := models.Image{
					ImageID:  databaseImage.ImageID,
					OHPostID: databaseImage.OHPostID,
					UserID:   databaseImage.UserID,
					Encoding: databaseImage.Base64Encode,
					XCoord:   databaseImage.XCoord,
					YCoord:   databaseImage.YCoord,
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

func GetMultipleImagesResponse(retrievedImages []models.Image) responses.MultipleImagesResponse {
	return responses.MultipleImagesResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: map[string]interface{}{
			"data": retrievedImages,
		},
	}
}
