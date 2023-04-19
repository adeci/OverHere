package ohpost_controller1

import (
	"OverHere/server/models"
	"OverHere/server/responses"
	"OverHere/server/services/database"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOHPostsByCoordBoundary() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cancel if request isn't processed in 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		//Get coordinates from route
		//:topleftcoord/:botrightcoord

		topLeftXCoord, conversionErr := strconv.ParseFloat(c.Param("topleftXcoord"), 64)
		if conversionErr != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestOHPostResponse(conversionErr.Error()),
			)
			cancel()
			return
		}

		topLeftYCoord, conversionErr := strconv.ParseFloat(c.Param("topleftYcoord"), 64)
		if conversionErr != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestOHPostResponse(conversionErr.Error()),
			)
			cancel()
			return
		}

		bottomRightXCoord, conversionErr := strconv.ParseFloat(c.Param("botrightXcoord"), 64)
		if conversionErr != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestOHPostResponse(conversionErr.Error()),
			)
			cancel()
			return
		}

		bottomRightYCoord, conversionErr := strconv.ParseFloat(c.Param("botrightYcoord"), 64)
		if conversionErr != nil {
			c.JSON(
				http.StatusBadRequest,
				BadRequestOHPostResponse(conversionErr.Error()),
			)
			cancel()
			return
		}

		defer cancel()

		fmt.Print("Getting by boundary: " + "tlc: " + fmt.Sprintf("%f", topLeftXCoord) + ", " + fmt.Sprintf("%f", topLeftYCoord) + ": brc: " + fmt.Sprintf("%f", bottomRightXCoord) + ", " + fmt.Sprintf("%f", bottomRightYCoord))

		allDatabaseOHPosts, err := database.GetOHPost_All()

		if err != nil {
			fmt.Println(err)
			cancel()
			return
		}

		ohpostsInsideBounds := []models.OHPost{}

		for _, element := range allDatabaseOHPosts {
			insideX := element.XCoord < topLeftXCoord && element.XCoord > bottomRightXCoord
			insideY := element.YCoord > topLeftYCoord && element.YCoord < bottomRightYCoord

			if insideX && insideY {
				ohpost := models.OHPost{
					OHPostID:  element.OHPostID,
					UserID:    element.UserID,
					Tag:       element.Tag,
					Caption:   element.Description,
					AvgXCoord: element.XCoord,
					AvgYCoord: element.YCoord,
				}

				ohpost = ReplaceXYCoordsIfInvalid(ohpost)
				ohpostsInsideBounds = append(ohpostsInsideBounds, ohpost)
			}
		}

		if err == nil {
			c.JSON(http.StatusOK, GetMultipleOHPostsResponse(ohpostsInsideBounds))
		} else {
			c.JSON(http.StatusBadRequest, BadRequestOHPostResponse(err.Error()))
		}
	}
}

func GetMultipleOHPostsResponse(retrievedOHPosts []models.OHPost) responses.MultipleOHPostResponse {
	return responses.MultipleOHPostResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: map[string]interface{}{
			"data": retrievedOHPosts,
		},
	}
}
