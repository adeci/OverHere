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

		var postOHPostRequest models.OHPost

		//Validate the request body
		if err := c.BindJSON(&postOHPostRequest); err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestOHPostResponse(err.Error()),
			)
			return
		}

		//Use the validator library to validate required fields
		if validationErr := helpers.Validate(&postOHPostRequest); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestOHPostResponse(validationErr.Error()),
			)
			cancel()
			return
		}

		//Logic
		//Averaging
		databaseOHPost, err := database.PostOHPost(postOHPostRequest.UserID, postOHPostRequest.Caption, 0.0, 0.0, postOHPostRequest.Tag)
		newOHPost := models.OHPost{
			OHPostID:  databaseOHPost.OHPostID,
			UserID:    databaseOHPost.UserID,
			Tag:       databaseOHPost.Tag,
			Caption:   databaseOHPost.Description,
			AvgXCoord: 0.0,
			AvgYCoord: 0.0,
		}

		fmt.Print(newOHPost)

		if err == nil {
			//Successful Response
			c.JSON(
				http.StatusCreated,
				PostedOHPostResponse(newOHPost),
			)
		} else {
			c.JSON(http.StatusBadRequest, BadRequestOHPostResponse(err.Error()))
		}
	}
}

func PostedOHPostResponse(newOHPost models.OHPost) responses.OHPostResponse {
	return responses.OHPostResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data: map[string]interface{}{
			"data": newOHPost,
		},
	}
}
