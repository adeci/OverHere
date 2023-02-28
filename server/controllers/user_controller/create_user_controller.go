package user_controller

//Followed tutorial for setup: https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m

import (
	"OverHere/server/models/user_model"
	"OverHere/server/responses/user_response"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user user_model.User

		//Validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestUserResponse(err.Error()),
			)
			return
		}

		//Use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestUserResponse(validationErr.Error()),
			)
			return
		}

		//Logic
		newUser := user_model.User{
			ObjectID: primitive.NewObjectID(),
			UserId:   user.UserId,
			Username: user.Username,
		}

		fmt.Print(newUser)

		//Successful Response
		c.JSON(
			http.StatusCreated,
			CreatedUserResponse(newUser),
		)
	}
}

func CreatedUserResponse(newUser user_model.User) user_response.UserResponse {
	return user_response.UserResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data: map[string]interface{}{
			"data": newUser,
		},
	}
}
