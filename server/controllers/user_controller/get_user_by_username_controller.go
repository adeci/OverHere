package user_controller

import (
	"OverHere/server/models"
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserByUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if request isn't processed in 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userID := c.Param("username")
		defer cancel()

		fmt.Print("Getting user by username: " + userID)

		retrievedUser, err := database.GetUser_Username(userID)

		user := models.User{
			UserID:   retrievedUser.UserID,
			Username: retrievedUser.Username,
		}

		if err == nil {
			c.JSON(http.StatusOK, GetUserResponse(user))
		} else {
			c.JSON(http.StatusBadRequest, BadRequestUserResponse(err.Error()))
		}
	}
}
