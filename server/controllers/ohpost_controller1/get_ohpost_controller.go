package ohpost_controller1

import (
	"OverHere/server/models"
	"OverHere/server/responses"
	"OverHere/server/services/database"
	"context"
	"fmt"
	"math"
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

		if err != nil {
			c.JSON(http.StatusBadRequest, BadRequestOHPostResponse(err.Error()))
			cancel()
			return
		}

		ohpost := models.OHPost{
			OHPostID:  retrievedOHPost.OHPostID,
			UserID:    retrievedOHPost.UserID,
			Tag:       "Blank tag",
			Caption:   retrievedOHPost.Description,
			AvgXCoord: retrievedOHPost.XCoord,
			AvgYCoord: retrievedOHPost.YCoord,
		}

		ohpost = ReplaceXYCoordsIfInvalid(ohpost)

		if err == nil {
			c.JSON(http.StatusOK, GetOHPostResponse(ohpost))
		} else {
			c.JSON(http.StatusBadRequest, BadRequestOHPostResponse(err.Error()))
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

func ReplaceXYCoordsIfInvalid(ohpost models.OHPost) models.OHPost {
	//Cannot display NaN
	if math.IsNaN(ohpost.AvgXCoord) {
		ohpost.AvgXCoord = 0
	}
	if math.IsNaN(ohpost.AvgYCoord) {
		ohpost.AvgYCoord = 0
	}

	return ohpost
}
