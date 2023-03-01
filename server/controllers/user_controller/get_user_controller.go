package user_controller

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		fmt.Print("Getting user")

		/*
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
			if validationErr := helpers.Validate(&user); validationErr != nil {
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
			) */
	}
}
