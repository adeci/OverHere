package image_controller

import (
	"context"
	"fmt"
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
