package ohpost_controller1

//Followed tutorial for setup: https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m

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

func PostOHPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

		//Logic
		//Averaging
		databaseOHPost, err := database.PostOHPost(ohpost.UserID, ohpost.Caption, 90.1, 80.1)
		newOHPost := models.OHPost{
			OHPostID:  databaseOHPost.OHPostID,
			UserID:    databaseOHPost.UserID,
			Tag:       "Blank tag",
			Caption:   databaseOHPost.Description,
			AvgXCoord: 90.1,
			AvgYCoord: 80.1,
		}

		fmt.Print(newOHPost)

		if err == nil {
			//Successful Response
			c.JSON(
				http.StatusCreated,
				CreatedOHPostResponse(newOHPost),
			)
		} else {
			c.JSON(http.StatusBadRequest, BadRequestOHPostResponse(err.Error()))
		}
	}
}

func CreatedOHPostResponse(newOHPost models.OHPost) responses.OHPostResponse {
	return responses.OHPostResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data: map[string]interface{}{
			"data": newOHPost,
		},
	}
}
