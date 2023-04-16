package ohpost_controller1

import (
	"OverHere/server/responses"
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DeleteOHPostsByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if request isn't processed in 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userID := c.Param("userid")
		defer cancel()

		fmt.Print("Deleting by userID " + userID)

		err := database.DeleteOHPost_UserID(userID)

		if err == nil {
			c.JSON(http.StatusOK, DeleteMultipleOHPostResponse())
		} else {
			c.JSON(http.StatusBadRequest, BadRequestOHPostResponse(err.Error()))
		}
	}
}

func DeleteMultipleOHPostResponse() responses.OHPostResponse {
	return responses.OHPostResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    map[string]interface{}{},
	}
}
