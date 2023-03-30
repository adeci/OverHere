package ohpost_controller1

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

func PutOHPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if request isn't processed in 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ohpostID := c.Param("ohpostid")
		defer cancel()

		var ohpost models.OHPost

		//Validate the request body
		if err := c.BindJSON(&ohpost); err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestOHPostResponse(err.Error()),
			)
			return
		}

		//Use the validator library to validate required fields
		if validationErr := helpers.Validate(&ohpost); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestOHPostResponse(validationErr.Error()),
			)
			return
		}

		fmt.Print("Getting user: " + ohpostID)

		databaseOHPostToPut := database.OHPostObject{
			OHPostID:    ohpostID,
			UserID:      ohpost.UserID,
			Description: ohpost.Caption,
			XCoord:      ohpost.AvgXCoord,
			YCoord:      ohpost.AvgYCoord,
		}

		database.PutOHPost(databaseOHPostToPut)

		retrievedOHPost, get_err := database.GetOHPost_OHPostID(ohpostID)

		databaseOHPostChanged := models.OHPost{
			OHPostID:  retrievedOHPost.OHPostID,
			UserID:    retrievedOHPost.UserID,
			Caption:   retrievedOHPost.Description,
			AvgXCoord: retrievedOHPost.XCoord,
			AvgYCoord: retrievedOHPost.YCoord,
		}

		if get_err == nil {
			c.JSON(http.StatusOK, PutOHPostResponse(databaseOHPostChanged))
		} else {
			c.JSON(http.StatusBadRequest, BadRequestOHPostResponse(get_err.Error()))
		}
	}
}

func PutOHPostResponse(retrievedOHPost models.OHPost) responses.UserResponse {
	return responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: map[string]interface{}{
			"data": retrievedOHPost,
		},
	}
}
