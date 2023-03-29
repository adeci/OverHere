package user_controller

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

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

		//Logic
		databaseUser, _ := database.GetUser_UserID(user.UserID)

		newUser := models.User{
			UserID:   databaseUser.UserID,
			Username: databaseUser.Username,
		}

		fmt.Print(newUser)

		//Successful Response
		c.JSON(
			http.StatusCreated,
			CreatedUserResponse(newUser),
		)
	}
}

func CreatedUserResponse(newUser models.User) responses.UserResponse {
	return responses.UserResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data: map[string]interface{}{
			"data": newUser,
		},
	}
}
