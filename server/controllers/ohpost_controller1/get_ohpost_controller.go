package ohpost_controller1

import (
	"OverHere/server/models"
	"OverHere/server/responses"
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOHPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if request isn't processed in 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ohpostID := c.Param("ohpostid")
		defer cancel()

		fmt.Print("Getting user: " + ohpostID)

		retrievedOHPost, err := database.GetOHPost_OHPostID(ohpostID)

		user := models.OHPost{
			OHPostID: retrievedOHPost.OHPostID,
			UserID:   retrievedOHPost.UserID,
			Tag:      "Blank tag",
			Caption:  retrievedOHPost.Description,
		}

		if err == nil {
			c.JSON(http.StatusOK, GetOHPostResponse(user))
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	}
}

func GetOHPostResponse(retrievedOHPost models.OHPost) responses.OHPostResponse {
	return responses.OHPostResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: map[string]interface{}{
			"data": retrievedOHPost,
		},
	}
}
