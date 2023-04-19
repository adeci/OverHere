package user_controller

import (
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DeleteUserByUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if request isn't processed in 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		username := c.Param("username")
		defer cancel()

		fmt.Print("Deleting user by username: " + username)

		databaseUser, err := database.GetUser_Username(username)

		if err != nil {
			c.JSON(http.StatusBadRequest, BadRequestUserResponse(err.Error()))
		}

		err = database.DeleteUser_Username(username)

		//Deleting user removes their OHPost and Images
		database.DeleteOHPost_UserID(databaseUser.UserID)
		database.DeleteImage_UserID(databaseUser.UserID)

		if err == nil {
			c.JSON(http.StatusOK, DeleteUserResponse())
		} else {
			c.JSON(http.StatusBadRequest, BadRequestUserResponse(err.Error()))
		}
	}
}
