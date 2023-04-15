package ohpost_controller1

import (
	"OverHere/server/models"
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOHPostsByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if request isn't processed in 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		//Get coordinates from route
		//:topleftcoord/:botrightcoord

		userID := c.Param("userid")
		defer cancel()

		allDatabaseOHPosts, err := database.GetOHPost_All()

		if err != nil {
			fmt.Println(err)
			cancel()
			return
		}

		ohpostsInsideBounds := []models.OHPost{}

		for _, element := range allDatabaseOHPosts {
			if userID == element.UserID {
				ohpost := models.OHPost{
					OHPostID:  element.OHPostID,
					UserID:    element.UserID,
					Tag:       "Blank tag",
					Caption:   element.Description,
					AvgXCoord: element.XCoord,
					AvgYCoord: element.YCoord,
				}

				ohpostsInsideBounds = append(ohpostsInsideBounds, ohpost)
			}
		}

		if err == nil {
			c.JSON(http.StatusOK, GetMultipleOHPostsResponse(ohpostsInsideBounds))
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	}
}
