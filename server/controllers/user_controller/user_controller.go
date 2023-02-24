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
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user user_model.User

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest,
				user_response.UserResponse{
					Status:  http.StatusBadRequest,
					Message: "error",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				user_response.UserResponse{
					Status:  http.StatusBadRequest,
					Message: "error",
					Data: map[string]interface{}{
						"data": validationErr.Error(),
					},
				},
			)
			return
		}

		newUser := user_model.User{
			Id:       primitive.NewObjectID(),
			Name:     user.Name,
			Location: user.Location,
			Title:    user.Title,
		}

		fmt.Print(newUser)

		c.JSON(http.StatusCreated, user_response.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": newUser}})
	}
}
