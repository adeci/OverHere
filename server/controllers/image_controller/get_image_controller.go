package image_controller

import (
	"OverHere/server/models/user_model"
	"OverHere/server/responses/user_response"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if enough time passes.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		fmt.Print("Creating image")

		/*
			var image := image_model.Image{

			}

			c.JSON(
				http.StatusFound,
				user_response.ImageResponse{
					Status:  http.StatusFound,
					Message: "success",
					Data: map[string]interface{}{
						"data": newUser,
					},
				},
			)
		*/
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
