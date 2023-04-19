package user_controller

import (
	"OverHere/server/responses"
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if request isn't processed in 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userID := c.Param("userid")
		defer cancel()

		fmt.Print("Getting user: " + userID)

		err := database.DeleteUser_UserID(userID)

		//Deleting user removes their OHPost and Images

		if err == nil {
			c.JSON(http.StatusOK, DeleteUserResponse())
		} else {
			c.JSON(http.StatusBadRequest, BadRequestUserResponse(err.Error()))
		}
	}
}

func DeleteUserResponse() responses.UserResponse {
	return responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    map[string]interface{}{},
	}
}
