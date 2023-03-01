package user_controller

import (
	"OverHere/server/models"
	"OverHere/server/responses"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if request isn't processed in 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userID := c.Param("userid")
		defer cancel()

		fmt.Print("Getting user: " + userID)

		user := models.User{
			ObjectID: primitive.NewObjectID(),
			UserID:   userID,
			Username: "Test",
		}

		c.JSON(http.StatusOK, GetUserResponse(user))
	}
}

func GetUserResponse(retrievedUser models.User) responses.UserResponse {
	return responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: map[string]interface{}{
			"data": retrievedUser,
		},
	}
}
