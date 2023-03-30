package user_controller

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

func PutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if request isn't processed in 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userID := c.Param("userid")
		defer cancel()

		var user models.User

		//Validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestUserResponse(err.Error()),
			)
			return
		}

		//Use the validator library to validate required fields
		if validationErr := helpers.Validate(&user); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestUserResponse(validationErr.Error()),
			)
			return
		}

		fmt.Print("Getting user: " + userID)

		put_err := database.PutUser(userID, user.Username)

		retrievedUser, get_err := database.GetUser_UserID(userID)

		databaseUserChanged := models.User{
			UserID:   retrievedUser.UserID,
			Username: retrievedUser.Username,
		}

		if put_err == nil && get_err == nil {
			c.JSON(http.StatusOK, PutUserResponse(databaseUserChanged))
		} else if put_err != nil {
			c.JSON(http.StatusBadRequest, BadRequestUserResponse(put_err.Error()))
		} else {
			c.JSON(http.StatusBadRequest, BadRequestUserResponse(get_err.Error()))
		}
	}
}

func PutUserResponse(retrievedUser models.User) responses.UserResponse {
	return responses.UserResponse{
		Status:  http.StatusAccepted,
		Message: "success",
		Data: map[string]interface{}{
			"data": retrievedUser,
		},
	}
}
