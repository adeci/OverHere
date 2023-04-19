package ohpost_controller1

import (
	"OverHere/server/models"
	"OverHere/server/services/database"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOHPostsByTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if request isn't processed in 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		tag := c.Param("tag")
		defer cancel()

		allDatabaseOHPosts, err := database.GetOHPost_All()

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestOHPostResponse(err.Error()),
			)
			cancel()
			return
		}

		matchingOHPosts := []models.OHPost{}

		for _, databaseOHPost := range allDatabaseOHPosts {
			if tag == databaseOHPost.Tag {
				ohpost := models.OHPost{
					OHPostID:  databaseOHPost.OHPostID,
					UserID:    databaseOHPost.UserID,
					Tag:       databaseOHPost.Tag,
					Caption:   databaseOHPost.Description,
					AvgXCoord: databaseOHPost.XCoord,
					AvgYCoord: databaseOHPost.YCoord,
				}

				ohpost = ReplaceXYCoordsIfInvalid(ohpost)
				matchingOHPosts = append(matchingOHPosts, ohpost)
			}
		}

		if err == nil {
			c.JSON(http.StatusOK, GetMultipleOHPostsResponse(matchingOHPosts))
		} else {
			c.JSON(http.StatusBadRequest, BadRequestOHPostResponse(err.Error()))
		}
	}
}
